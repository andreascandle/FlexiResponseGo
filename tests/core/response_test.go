package core_test

import (
	"testing"

	"github.com/andreascandle/FlexiResponseGo/core"
	"github.com/andreascandle/FlexiResponseGo/utils"
	"github.com/stretchr/testify/assert"
)

func TestNewSuccessResponse(t *testing.T) {
	resp := core.NewSuccessResponse("trace-123", "Success", map[string]string{"key": "value"})
	assert.Equal(t, "success", resp.Status)
	assert.Equal(t, "Success", resp.Message)
	assert.Equal(t, "trace-123", resp.TraceID)
	assert.Contains(t, resp.Metadata, "timestamp")
}

func TestNewErrorResponse(t *testing.T) {
	resp := core.NewErrorResponse("trace-123", "Error occurred", "Something went wrong")
	assert.Equal(t, "error", resp.Status)
	assert.Equal(t, "Error occurred", resp.Message)
	assert.Equal(t, "trace-123", resp.TraceID)
	assert.Equal(t, "Something went wrong", resp.Error)
}

func TestGenerateTraceID(t *testing.T) {
	traceID := utils.GenerateTraceID(16)
	assert.Equal(t, 16, len(traceID))
}
