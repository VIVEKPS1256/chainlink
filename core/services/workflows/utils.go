package workflows

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/attribute"
	"reflect"
)

const WorkflowID = "WorkflowID"
const WorkflowExecutionID = "WorkflowExecutionID"

type workflowContextKey struct{}

var keystoneContextKey = workflowContextKey{}

type KeystoneWorkflowLabels struct {
	WorkflowExecutionID string
	WorkflowID          string
}

var OrderedKeystoneLabels = []string{WorkflowID, WorkflowExecutionID}

var OrderedKeystoneLabelsMap = make(map[string]interface{})

func init() {
	for _, label := range OrderedKeystoneLabels {
		OrderedKeystoneLabelsMap[label] = interface{}(0)
	}
}

func (k *KeystoneWorkflowLabels) ToMap() map[string]string {
	labels := make(map[string]string)

	labels[WorkflowID] = k.WorkflowID
	labels[WorkflowExecutionID] = k.WorkflowExecutionID

	return labels
}

func (k *KeystoneWorkflowLabels) ToOtelAttributes() []attribute.KeyValue {
	return []attribute.KeyValue{
		attribute.String(WorkflowID, k.WorkflowID),
		attribute.String(WorkflowExecutionID, k.WorkflowExecutionID),
	}
}

// GetKeystoneLabelsFromContext extracts the KeystoneWorkflowLabels struct set on the
// unexported keystoneContextKey. Call NewKeystoneContext first before usage -
// if the key is unset or the value is not of the expected type GetKeystoneLabelsFromContext will error.
func GetKeystoneLabelsFromContext(ctx context.Context) (KeystoneWorkflowLabels, error) {
	curLabelsAny := ctx.Value(keystoneContextKey)
	curLabels, ok := curLabelsAny.(KeystoneWorkflowLabels)
	if !ok {
		return KeystoneWorkflowLabels{}, fmt.Errorf("context value with keystoneContextKey is not of type KeystoneWorkflowLabels")
	}

	return curLabels, nil
}

// NewKeystoneContext returns a context with the keystoneContextKey loaded. This enables
// the context to be consumed by GetKeystoneLabelsFromContext and KeystoneContextWithLabel.
// labels should not be nil.
func NewKeystoneContext(ctx context.Context, labels KeystoneWorkflowLabels) context.Context {
	return context.WithValue(ctx, keystoneContextKey, labels)
}

// KeystoneContextWithLabel extracts the Keystone Labels set on the passed in immutable context,
// sets the new desired label if valid, and then returns a new context with the updated labels
func KeystoneContextWithLabel(ctx context.Context, key string, value string) (context.Context, error) {
	curLabels, err := GetKeystoneLabelsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if OrderedKeystoneLabelsMap[key] == nil {
		return nil, fmt.Errorf("key %v is not a valid keystone label", key)
	}

	reflectedLabels := reflect.ValueOf(&curLabels).Elem()
	reflectedLabels.FieldByName(key).SetString(value)

	newLabels := reflectedLabels.Interface().(KeystoneWorkflowLabels)
	return context.WithValue(ctx, keystoneContextKey, newLabels), nil
}

func composeLabeledMsg(ctx context.Context, msg string) (string, error) {
	structLabels, err := GetKeystoneLabelsFromContext(ctx)
	if err != nil {
		return "", fmt.Errorf("composing labeled message failed: %w", err)
	}

	labels := structLabels.ToMap()

	// Populate labeled message in reverse
	numLabels := len(OrderedKeystoneLabels)
	for i := range numLabels {
		msg = fmt.Sprintf("%v.%v", labels[OrderedKeystoneLabels[numLabels-1-i]], msg)
	}

	return msg, nil
}

func getOtelAttributesFromCtx(ctx context.Context) ([]attribute.KeyValue, error) {
	labelsStruct, err := GetKeystoneLabelsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	otelLabels := labelsStruct.ToOtelAttributes()
	return otelLabels, nil
}
