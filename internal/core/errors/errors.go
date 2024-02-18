package errors

import "errors"

var (
	// ErrNotFound indicates that the requested resource was not found.
	ErrNotFound = errors.New("not found")

	// ErrUnauthorized indicates that the request lacks valid authentication credentials.
	ErrUnauthorized = errors.New("unauthorized")

	// ErrForbidden indicates that the request is valid but the user lacks necessary permissions.
	ErrForbidden = errors.New("forbidden")

	// ErrBadRequest indicates that the request is invalid or malformed.
	ErrBadRequest = errors.New("bad request")

	// ErrInternalServer indicates that an internal server error occurred.
	ErrInternalServer = errors.New("internal server error")
)
