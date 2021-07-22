package web

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/services/chainlink"
)

type ReplayController struct {
	App chainlink.Application
}

// ReplayFromBlock causes the node to process blocks again from the given block number
// Example:
//  "<application>/v2/replay_from_block/:number"
func (bdc *ReplayController) ReplayFromBlock(c *gin.Context) {

	if c.Param("number") == "" {
		jsonAPIError(c, http.StatusUnprocessableEntity, errors.New("missing 'number' parameter"))
		return
	}

	blockNumber, err := strconv.ParseInt(c.Param("number"), 10, 0)
	if err != nil {
		jsonAPIError(c, http.StatusUnprocessableEntity, err)
		return
	}
	if blockNumber < 0 {
		jsonAPIError(c, http.StatusUnprocessableEntity, errors.Errorf("block number cannot be negative: %v", blockNumber))
		return
	}
	if err := bdc.App.ReplayFromBlock(uint64(blockNumber)); err != nil {
		jsonAPIError(c, http.StatusInternalServerError, err)
		return
	}

	response := ReplayResponse{
		Message: "Replay started",
	}
	jsonAPIResponse(c, &response, "response")
}

type ReplayResponse struct {
	Message string `json:"message"`
}

// GetID returns the jsonapi ID.
func (s ReplayResponse) GetID() string {
	return "replayID"
}

// GetName returns the collection name for jsonapi.
func (ReplayResponse) GetName() string {
	return "replay"
}

// SetID is used to conform to the UnmarshallIdentifier interface for
// deserializing from jsonapi documents.
func (*ReplayResponse) SetID(string) error {
	return nil
}
