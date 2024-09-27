// Package errors provides error codes for the application.
package errors

const (
	// CodeInternalServerError is a generic error message returned when an unexpected error occurs.
	CodeInternalServerError = "INTERNAL_SERVER_ERROR"
	// CodeInvalidRequest is a generic error message returned when the request is invalid.
	CodeInvalidRequest = "INVALID_REQUEST"
	// CodeNotFound is a generic error message returned when the requested resource is not found.
	CodeNotFound = "NOT_FOUND"
	// CodeBadRequest is a generic error message returned when the request is bad.
	CodeBadRequest = "BAD_REQUEST"
)
