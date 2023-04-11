package errors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"runtime"
)

//inspiration: https://middlemost.com/failure-is-your-domain/

// interface guards
var _ error = (*Error)(nil)

// Application error codes.
const (
	ECONFLICT       = "conflict"     // action cannot be performed
	EINTERNAL       = "internal"     // internal error
	EBADREQUEST     = "bad_request"  // validation failed
	ENOTFOUND       = "not_found"    // entity does not exist
	EUNAUTHORIZED   = "unauthorized" // operation unauthorized
	EFORBIDDEN      = "forbidden"    //operation forbidden
	EEXPECTED       = "expected"     //expected error that don't need to be logged
	ETIMEOUT        = "timeout"
	ETOOMANYREQUEST = "too_many_request"
)

// Error defines a standard application error.
type Error struct {
	Code       string    // Machine-readable error code
	Message    string    // Human-readable message
	Op         string    // Logical operation and nested error
	Err        error     // Embedded error
	Detail     []byte    // JSON encoded data
	stacktrace []uintptr // stacktrace pointers for reporting error
}

// NewError creates a new error with the given code and message.
func NewError(e *Error) *Error {
	pcs := make([]uintptr, 100)
	n := runtime.Callers(2, pcs)
	e.stacktrace = pcs[:n]

	return e
}

// ErrorCode returns the code of the root error, if available.
// Otherwise, returns EINTERNAL.
func ErrorCode(err error) string {
	if err == nil {
		return ""
	}
	e, ok := err.(*Error)
	if ok && e.Code != "" {
		return e.Code
	}
	if ok && e.Err != nil {
		return ErrorCode(e.Err)
	}

	return EINTERNAL
}

// ErrorMessage returns the human-readable message of the error, if available.
// Otherwise, returns a generic error message.
func ErrorMessage(err error) string {
	if err == nil {
		return ""
	}
	e, ok := err.(*Error)
	if ok && e.Message != "" {
		return e.Message
	}
	if ok && e.Err != nil {
		return ErrorMessage(e.Err)
	}

	return "An internal error has occurred. Please contact technical support."
}

// StackTrace returns the stacktrace pointers for reporting error.
func (e *Error) StackTrace() []uintptr {
	var stacktrace []uintptr
	stacktrace = append(stacktrace, e.stacktrace...)
	if e.Err != nil {
		err, ok := e.Err.(*Error)
		if ok {
			stacktrace = append(stacktrace, err.StackTrace()...)
		}
	}
	return stacktrace
}

// Error returns the string representation of the error message.
func (e *Error) Error() string {
	var buf bytes.Buffer

	// Print the current operation in our stack, if any.
	if e.Op != "" {
		fmt.Fprintf(&buf, "%s: ", e.Op)
	}

	// If wrapping an error, print its Error() message.
	// Otherwise, print the error code & message.
	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.Code != "" {
			fmt.Fprintf(&buf, "<%s> ", e.Code)
		}
		buf.WriteString(e.Message)
	}
	return buf.String()
}

// DetailedError returns the detailed string representation of the error message.
func (e *Error) DetailedError() string {
	return walkError("", e)
}

func walkError(prefix string, e *Error) string {
	var buf bytes.Buffer
	if e.Op != "" {
		fmt.Fprintf(&buf, "%sOp: %s\n", prefix, e.Op)
	}
	if e.Code != "" {
		fmt.Fprintf(&buf, "%sCode: %s\n", prefix, e.Code)
	}
	if e.Message != "" {
		fmt.Fprintf(&buf, "%sMessage: %s\n", prefix, e.Message)
	}
	if string(e.Detail) != "" {
		switch json.Valid(e.Detail) {
		case true:
			fmt.Fprintf(&buf, "%sDetail:\n%s", prefix, prefix)
			json.Indent(&buf, e.Detail, prefix, "\t")
		default:
			fmt.Fprintf(&buf, "%sDetail: %s\n", prefix, e.Detail)
		}
	}
	if e.Err != nil {
		embedErr, ok := e.Err.(*Error)
		switch ok {
		case true:
			fmt.Fprintf(&buf, "%sErr:\n%s", prefix, walkError(fmt.Sprintf("%s\t", prefix), embedErr))
		default:
			fmt.Fprintf(&buf, "%sErr: %s\n", prefix, e.Err.Error())
		}
	}
	return buf.String()
}
