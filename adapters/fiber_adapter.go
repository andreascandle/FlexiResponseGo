package adapters

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// FiberSuccessResponse sends a success response in Fiber with logging.
func FiberSuccessResponse(c *fiber.Ctx, message string, data interface{}) error {
	start := time.Now()
	traceID := GetOrGenerateTraceID(c.GetReqHeaders())
	LogRequest(c.Method(), c.Path(), traceID, c.GetReqHeaders())

	resp := GenerateSuccessResponse(traceID, message, data)
	err := c.Status(fiber.StatusOK).JSON(resp)

	LogResponse(c.Method(), c.Path(), traceID, fiber.StatusOK, time.Since(start))
	return err
}

// FiberErrorResponse sends an error response in Fiber with logging.
func FiberErrorResponse(c *fiber.Ctx, statusCode int, message, errorDetail string) error {
	start := time.Now()
	traceID := GetOrGenerateTraceID(c.GetReqHeaders())
	LogRequest(c.Method(), c.Path(), traceID, c.GetReqHeaders())

	resp := GenerateErrorResponse(traceID, message, errorDetail)
	err := c.Status(statusCode).JSON(resp)

	LogResponse(c.Method(), c.Path(), traceID, statusCode, time.Since(start))
	return err
}
