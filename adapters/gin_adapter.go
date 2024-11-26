package adapters

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GinSuccessResponse sends a success response in Gin with logging.
func GinSuccessResponse(c *gin.Context, message string, data interface{}) {
	start := time.Now()
	traceID := GetOrGenerateTraceID(c.Request.Header)
	LogRequest(c.Request.Method, c.Request.URL.Path, traceID, c.Request.Header)

	resp := GenerateSuccessResponse(traceID, message, data)
	c.JSON(http.StatusOK, resp)

	LogResponse(c.Request.Method, c.Request.URL.Path, traceID, http.StatusOK, time.Since(start))
}

// GinErrorResponse sends an error response in Gin with logging.
func GinErrorResponse(c *gin.Context, statusCode int, message, errorDetail string) {
	start := time.Now()
	traceID := GetOrGenerateTraceID(c.Request.Header)
	LogRequest(c.Request.Method, c.Request.URL.Path, traceID, c.Request.Header)

	resp := GenerateErrorResponse(traceID, message, errorDetail)
	c.JSON(statusCode, resp)

	LogResponse(c.Request.Method, c.Request.URL.Path, traceID, statusCode, time.Since(start))
}
