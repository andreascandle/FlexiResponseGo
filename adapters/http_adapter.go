package adapters

import (
	"net/http"
	"time"
)

// HTTPSuccessResponse sends a success response for net/http with logging.
func HTTPSuccessResponse(w http.ResponseWriter, r *http.Request, message string, data interface{}) {
	start := time.Now()
	traceID := GetOrGenerateTraceID(r.Header)
	LogRequest(r.Method, r.URL.Path, traceID, r.Header)

	resp := GenerateSuccessResponse(traceID, message, data)
	WriteJSONResponse(w, http.StatusOK, resp)

	LogResponse(r.Method, r.URL.Path, traceID, http.StatusOK, time.Since(start))
}

// HTTPErrorResponse sends an error response for net/http with logging.
func HTTPErrorResponse(w http.ResponseWriter, r *http.Request, statusCode int, message, errorDetail string) {
	start := time.Now()
	traceID := GetOrGenerateTraceID(r.Header)
	LogRequest(r.Method, r.URL.Path, traceID, r.Header)

	resp := GenerateErrorResponse(traceID, message, errorDetail)
	WriteJSONResponse(w, statusCode, resp)

	LogResponse(r.Method, r.URL.Path, traceID, statusCode, time.Since(start))
}
