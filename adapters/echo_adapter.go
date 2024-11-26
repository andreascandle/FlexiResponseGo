package adapters

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// EchoSuccessResponse sends a success response in Echo with logging.
func EchoSuccessResponse(c echo.Context, message string, data interface{}) error {
	start := time.Now()
	traceID := GetOrGenerateTraceID(c.Request().Header)
	LogRequest(c.Request().Method, c.Request().URL.Path, traceID, c.Request().Header)

	resp := GenerateSuccessResponse(traceID, message, data)
	err := c.JSON(http.StatusOK, resp)

	LogResponse(c.Request().Method, c.Request().URL.Path, traceID, http.StatusOK, time.Since(start))
	return err
}

// EchoErrorResponse sends an error response in Echo with logging.
func EchoErrorResponse(c echo.Context, statusCode int, message, errorDetail string) error {
	start := time.Now()
	traceID := GetOrGenerateTraceID(c.Request().Header)
	LogRequest(c.Request().Method, c.Request().URL.Path, traceID, c.Request().Header)

	resp := GenerateErrorResponse(traceID, message, errorDetail)
	err := c.JSON(statusCode, resp)

	LogResponse(c.Request().Method, c.Request().URL.Path, traceID, statusCode, time.Since(start))
	return err
}
