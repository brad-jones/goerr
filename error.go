package goerr

import (
	"fmt"
)

// Error is an error object that stores stack frame information,
// create new instances with `New()`.
type Error struct {
	message  string
	innerErr error
	caller   uintptr
}

// New is the constructor for the `Error` object.
//
// It will accept any value and convert it to an error using
// `fmt.Errorf("%v", value)` if need be and then set the
// resulting value as the innerErr.
func New(value interface{}) *Error {
	err, ok := value.(error)
	if !ok {
		err = fmt.Errorf("%+v", value)
	}
	return &Error{innerErr: err}
}

// Error implements the stdlib error interface.
func (g *Error) Error() string {
	if g.message == "" {
		return g.innerErr.Error()
	}
	return fmt.Sprintf("%s: %s", g.message, g.innerErr.Error())
}

// Unwrap implements the stdlib error interface.
func (g *Error) Unwrap() error {
	return g.innerErr
}

// Frame returns the stack frame object attached to this error.
func (g *Error) Frame() *StackFrame {
	return NewStackFrame(g.caller)
}
