package core

// ErrorCategory defines different categories of errors.
type ErrorCategory string

const (
	ClientError          ErrorCategory = "client_error"
	ServerError          ErrorCategory = "server_error"
	ValidationError      ErrorCategory = "validation_error"
	RateLimitError       ErrorCategory = "rate_limit_error"
	AuthenticationError  ErrorCategory = "authentication_error"
	AuthorizationError   ErrorCategory = "authorization_error"
	DatabaseError        ErrorCategory = "database_error"
	ExternalServiceError ErrorCategory = "external_service_error"
)

// APIError represents a structured error type.
type APIError struct {
	Category    ErrorCategory          `json:"category"`
	Code        int                    `json:"code"`
	Message     string                 `json:"message"`
	Details     string                 `json:"details"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	NestedError *APIError              `json:"nested_error,omitempty"`
}

// NewAPIError creates a new API error with optional metadata.
func NewAPIError(category ErrorCategory, code int, message, details string) APIError {
	return APIError{
		Category: category,
		Code:     code,
		Message:  message,
		Details:  sanitizeError(details),
	}
}

// WithDetails sets or updates the Details field of an APIError.
func (e APIError) WithDetails(details string) APIError {
	e.Details = sanitizeError(details)
	return e
}

// WithMetadata adds or updates metadata in an APIError.
func (e APIError) WithMetadata(key string, value interface{}) APIError {
	if e.Metadata == nil {
		e.Metadata = make(map[string]interface{})
	}
	e.Metadata[key] = value
	return e
}

// WithNestedError attaches a nested error for context.
func (e APIError) WithNestedError(nested APIError) APIError {
	e.NestedError = &nested
	return e
}

// IsClientError checks if the error is a client-side error.
func IsClientError(err APIError) bool {
	return err.Category == ClientError || err.Category == ValidationError || err.Category == RateLimitError
}

// IsServerError checks if the error is a server-side error.
func IsServerError(err APIError) bool {
	return err.Category == ServerError || err.Category == DatabaseError || err.Category == ExternalServiceError
}

// sanitizeError ensures sensitive details are not exposed in errors.
func sanitizeError(details string) string {
	if len(details) > 100 {
		return "An internal error occurred."
	}
	return details
}
