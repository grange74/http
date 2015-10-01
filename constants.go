package http

// Common HTTP Methods.
const (
	MethodGet    = "GET"
	MethodPost   = "POST"
	MethodDelete = "DELETE"
	MethodPut    = "PUT"
)

// Common Content Types.
const (
	ContentTypeJSON = "application/json"
	ContentTypeText = "text/plain"
	ContentTypeNone = ""
)

// Most common HTTP status codes.
const (
	StatusOK        = 200
	StatusCreated   = 201
	StatusNoContent = 204

	StatusNotModified = 304

	StatusBadRequest   = 400
	StatusUnauthorized = 401
	StatusForbidden    = 403
	StatusNotFound     = 404
	StatusConflict     = 409

	StatusInternalServerError = 500
)

// Defaults
const (
	DefaultTimeout = 30 // in seconds
)
