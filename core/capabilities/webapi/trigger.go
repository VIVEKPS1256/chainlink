package webapi

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	ethCommon "github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-common/pkg/types/core"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/gateway/api"
	"github.com/smartcontractkit/chainlink/v2/core/services/gateway/connector"
	"github.com/smartcontractkit/chainlink/v2/core/services/gateway/handlers/workflow"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
)

const defaultSendChannelBufferSize = 1000

type Response struct {
	Success      bool   `json:"success"`
	ErrorMessage string `json:"error_message,omitempty"`
}

type workflowConnectorHandler struct {
	services.StateMachine

	capabilities.CapabilityInfo
	connector connector.GatewayConnector
	lggr      logger.Logger
	mu        sync.Mutex
	triggers  map[string]chan capabilities.TriggerResponse
}

var _ capabilities.TriggerCapability = (*workflowConnectorHandler)(nil)
var _ services.Service = &workflowConnectorHandler{}

func NewTrigger(config string, registry core.CapabilitiesRegistry, connector connector.GatewayConnector, lggr logger.Logger) (job.ServiceCtx, error) {
	// TODO (CAPPL-22, CAPPL-24):
	//   - decode config
	//   - create an implementation of the capability API and add it to the Registry
	//   - create a handler and register it with Gateway Connector
	//   - manage trigger subscriptions
	//   - process incoming trigger events and related metadata

	handler := &workflowConnectorHandler{
		connector: connector,
		lggr:      lggr.Named("WorkflowConnectorHandler"),
	}

	// is this the right way to register with gateway connector?  Cron trigger doesn't do this.
	err := connector.AddHandler([]string{"add_workflow"}, handler)

	return handler, err
}

func (h *workflowConnectorHandler) HandleGatewayMessage(ctx context.Context, gatewayID string, msg *api.Message) {
	body := &msg.Body
	fromAddr := ethCommon.HexToAddress(body.Sender)
	// TODO: apply allowlist and rate-limiting
	h.lggr.Debugw("handling gateway request", "id", gatewayID, "method", body.Method, "address", fromAddr)

	switch body.Method {
	case workflow.MethodAddWorkflow:
		// TODO: add a new workflow spec and return success/failure
		// we need access to Job ORM or whatever CLO uses to fully launch a new spec
		h.lggr.Debugw("added workflow spec", "payload", string(body.Payload))
		response := Response{Success: true}
		h.sendResponse(ctx, gatewayID, body, response)
	default:
		h.lggr.Errorw("unsupported method", "id", gatewayID, "method", body.Method)
	}
}

// Register a new trigger
// Can register triggers before the service is actively scheduling
func (h *workflowConnectorHandler) RegisterTrigger(ctx context.Context, req capabilities.TriggerRegistrationRequest) (<-chan capabilities.TriggerResponse, error) {
	// There's no config to use and validate
	h.mu.Lock()
	defer h.mu.Unlock()
	_, ok := h.triggers[req.TriggerID]
	if ok {
		return nil, fmt.Errorf("triggerId %s already registered", req.TriggerID)
	}

	callbackCh := make(chan capabilities.TriggerResponse, defaultSendChannelBufferSize)
	h.triggers[req.TriggerID] = callbackCh

	return callbackCh, nil
}

func (h *workflowConnectorHandler) UnregisterTrigger(ctx context.Context, req capabilities.TriggerRegistrationRequest) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	trigger := h.triggers[req.TriggerID]

	// Close callback channel
	close(trigger)
	// Remove from triggers context
	delete(h.triggers, req.TriggerID)
	return nil
}

func (h *workflowConnectorHandler) Start(ctx context.Context) error {
	// how does the
	return h.StartOnce("GatewayConnectorServiceWrapper", func() error {
		return nil
	})
}
func (h *workflowConnectorHandler) Close() error {
	return nil
}

func (h *workflowConnectorHandler) Ready() error {
	return nil
}

func (h *workflowConnectorHandler) HealthReport() map[string]error {
	return map[string]error{h.Name(): nil}
}

func (h *workflowConnectorHandler) Name() string {
	return "WebAPITrigger"
}

func (h *workflowConnectorHandler) sendResponse(ctx context.Context, gatewayID string, requestBody *api.MessageBody, payload any) error {
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	msg := &api.Message{
		Body: api.MessageBody{
			MessageId: requestBody.MessageId,
			DonId:     requestBody.DonId,
			Method:    requestBody.Method,
			Receiver:  requestBody.Sender,
			Payload:   payloadJson,
		},
	}

	// How do we get the signerKey from the connector?
	// if err = msg.Sign(h.signerKey); err != nil {
	// 	return err
	// }
	return h.connector.SendToGateway(ctx, gatewayID, msg)
}
