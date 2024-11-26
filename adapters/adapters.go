package adapters

import (
	"net/http"
	"time"

	"github.com/andreascandle/FlexiResponseGo/core"
	"github.com/andreascandle/FlexiResponseGo/logger"
	"github.com/andreascandle/FlexiResponseGo/utils"
	"go.uber.org/zap"
)

// GetOrGenerateTraceID retrieves a trace ID from headers or generates a new one.
func GetOrGenerateTraceID(headers http.Header) string {
	traceID := headers.Get("X-Trace-ID")
	if traceID == "" {
		traceID = utils.GenerateTraceID(16)
		headers.Set("X-Trace-ID", traceID)
	}
	return traceID
}

// LogRequest logs incoming request details.
func LogRequest(method, path, traceID string, headers http.Header) {
	log := logger.GetLogger()
	log.Info("Incoming request",
		zap.String("trace_id", traceID),
		zap.String("method", method),
		zap.String("path", path),
		zap.Any("headers", headers),
	)
}

// LogResponse logs outgoing response details.
func LogResponse(method, path, traceID string, statusCode int, duration time.Duration) {
	log := logger.GetLogger()
	log.Info("Outgoing response",
		zap.String("trace_id", traceID),
		zap.String("method", method),
		zap.String("path", path),
		zap.Int("status_code", statusCode),
		zap.Duration("duration", duration),
	)
}

// GenerateSuccessResponse creates a standardized success response.
func GenerateSuccessResponse(traceID, message string, data interface{}) core.StandardResponse {
	return core.NewSuccessResponse(traceID, message, data)
}

// GenerateErrorResponse creates a standardized error response.
func GenerateErrorResponse(traceID, message, errorDetail string) core.StandardResponse {
	return core.NewErrorResponse(traceID, message, errorDetail)
}

// GenerateValidationErrorResponse creates a validation error response.
func GenerateValidationErrorResponse(traceID, message string, fieldErrors map[string]interface{}) core.StandardResponse {
	return core.NewValidationErrorResponse(traceID, message, fieldErrors)
}

// WriteJSONResponse writes the response to the client.
func WriteJSONResponse(w http.ResponseWriter, statusCode int, response core.StandardResponse) {
	core.WriteJSON(w, statusCode, response)
}
