package jayson

import (
	"fmt"
)

// A ErrorCode by JSON-RPC 2.0.
const (
	// ErrorCodeParse is parse error code.
	ErrorCodeParse int = -32700
	// ErrorCodeInvalidRequest is invalid request error code.
	ErrorCodeInvalidRequest int = -32600
	// ErrorCodeMethodNotFound is method not found error code.
	ErrorCodeMethodNotFound int = -32601
	// ErrorCodeInvalidParams is invalid params error code.
	ErrorCodeInvalidParams int = -32602
	// ErrorCodeInternal is internal error code.
	ErrorCodeInternal int = -32603
)

// An Error is a wrapper for a JSON interface value.
type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Error implements error interface.
func (e *Error) Error() string {
	return fmt.Sprintf("jsonrpc: code: %d, message: %s, data: %+v", e.Code, e.Message, e.Data)
}

// ErrParse returns parse error.
func ErrParse() *Error {
	return &Error{
		Code:    ErrorCodeParse,
		Message: "PARSE_ERROR",
	}
}

// ErrInvalidRequest returns invalid request error.
func ErrInvalidRequest() *Error {
	return &Error{
		Code:    ErrorCodeInvalidRequest,
		Message: "INVALID_REQUEST",
	}
}

// ErrMethodNotFound returns method not found error.
func ErrMethodNotFound() *Error {
	return &Error{
		Code:    ErrorCodeMethodNotFound,
		Message: "METHOD_NOT_FOUND",
	}
}

// ErrInvalidParams returns invalid params error.
func ErrInvalidParams() *Error {
	return &Error{
		Code:    ErrorCodeInvalidParams,
		Message: "INVALID_PARAMS",
	}
}

// ErrInternal returns internal error.
func ErrInternal() *Error {
	return &Error{
		Code:    ErrorCodeInternal,
		Message: "INTERNAL_ERROR",
	}
}
