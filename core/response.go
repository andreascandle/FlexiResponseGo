package core

import (
	"bytes"
	"net/http"
	"time"

	"github.com/andreascandle/FlexiResponseGo/config"
	"github.com/andreascandle/FlexiResponseGo/utils"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// StandardResponse defines the unified API response structure.
type StandardResponse struct {
	Status      string                 `json:"status"`
	Message     string                 `json:"message"`
	Data        interface{}            `json:"data,omitempty"`
	Error       string                 `json:"error,omitempty"`
	TraceID     string                 `json:"trace_id,omitempty"`
	FieldErrors map[string]interface{} `json:"field_errors,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// NewSuccessResponse creates a standardized success response.
func NewSuccessResponse(traceID, message string, data interface{}) StandardResponse {
	if traceID == "" {
		traceID = utils.GenerateTraceID(16)
	}
	return StandardResponse{
		Status:  "success",
		Message: localizeMessage(message),
		Data:    data,
		TraceID: traceID,
		Metadata: mergeMetadata(map[string]interface{}{
			"timestamp": time.Now().Format(time.RFC3339),
		}),
	}
}

// NewErrorResponse creates a standardized error response.
func NewErrorResponse(traceID, message, errorDetail string) StandardResponse {
	if traceID == "" {
		traceID = utils.GenerateTraceID(16)
	}
	return StandardResponse{
		Status:  "error",
		Message: localizeMessage(message),
		Error:   sanitizeError(errorDetail),
		TraceID: traceID,
		Metadata: mergeMetadata(map[string]interface{}{
			"timestamp": time.Now().Format(time.RFC3339),
		}),
	}
}

// NewValidationErrorResponse creates a response for validation errors.
func NewValidationErrorResponse(traceID, message string, fieldErrors map[string]interface{}) StandardResponse {
	if traceID == "" {
		traceID = utils.GenerateTraceID(16)
	}
	return StandardResponse{
		Status:      "error",
		Message:     localizeMessage(message),
		FieldErrors: fieldErrors,
		TraceID:     traceID,
		Metadata: mergeMetadata(map[string]interface{}{
			"timestamp": time.Now().Format(time.RFC3339),
		}),
	}
}

// WriteJSON sends a JSON response with optimal performance.
func WriteJSON(w http.ResponseWriter, statusCode int, resp StandardResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Trace-ID", resp.TraceID)
	w.WriteHeader(statusCode)

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(resp); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return err
	}

	_, err := w.Write(buf.Bytes())
	return err
}

// WriteErrorResponse writes an APIError to the response using StandardResponse.
func WriteErrorResponse(w http.ResponseWriter, statusCode int, traceID string, apiErr APIError) error {
	resp := StandardResponse{
		Status:  "error",
		Message: apiErr.Message,
		Error:   apiErr.Details,
		TraceID: traceID,
		Metadata: mergeMetadata(map[string]interface{}{
			"category":  apiErr.Category,
			"code":      apiErr.Code,
			"timestamp": time.Now().Format(time.RFC3339),
		}),
	}
	return WriteJSON(w, statusCode, resp)
}

// mergeMetadata combines global and local metadata dynamically.
func mergeMetadata(localMetadata map[string]interface{}) map[string]interface{} {
	conf := config.GetConfig()
	for k, v := range conf.GlobalMetadata {
		localMetadata[k] = v
	}
	return localMetadata
}

// localizeMessage returns a localized message if localization is enabled.
func localizeMessage(message string) string {
	conf := config.GetConfig()
	if conf.GlobalMetadata["enableLocalization"] == true {
		// Placeholder for localization logic. Extend with an i18n system or database.
		return message
	}
	return message
}
