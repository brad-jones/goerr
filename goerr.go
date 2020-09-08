package goerr

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
)

// PrintTrace will print the stack trace for the given error to STDERR.
//
// Included will be the error message, if the error is of type *goerr.Error then
// a stack trace will be generated and finally if the cause of the error can be
// marshalled into JSON it's values will be output as well.
//
// For example:
//  human friendly error message
//
//  {
//  	"optional": "context values"
//  }
//
//  the-pkg-name.theMethodName:/the/file/path/to/go/src/file:123
//  	if /the/file/path/to/go/src/file exists then this will be line 123
func PrintTrace(err error) {
	fmt.Fprint(os.Stderr, NewStackTrace(err).String())
}

// Trace will take any value, convert it into an Error object,
// if not already one and then save the stack trace information.
//
// It also accepts a variadic number of messages that will be
// prefixed to the error text it's self to provide additional
// context if required. These messages should be human friendly.
func Trace(value interface{}, messages ...string) *Error {
	return trace(1, value, messages...)
}

func trace(skip int, value interface{}, messages ...string) *Error {
	err, ok := value.(*Error)
	if !ok {
		err = New(value)
	} else {
		if err.callers == nil {
			err = &Error{innerErr: err}
		}
	}

	caller := skip + 1
	callers := []uintptr{}
	for {
		pc, _, _, ok := runtime.Caller(caller)
		if !ok {
			break
		}
		callers = append(callers, pc)
		caller++
	}

	if err.callers == nil {
		err.callers = callers
	}

	return &Error{
		innerErr: err,
		callers:  callers[1:],
		message:  strings.Join(messages, ": "),
	}
}

// Check will panic if err is not nill.
// It does the same tracing as the Trace function.
//
// YMMV - It mimics the goV2 check/handle proposal: https://bit.ly/354fRXv
func Check(err error, messages ...string) {
	if err != nil {
		panic(trace(1, err, messages...))
	}
}

// Handle will recover, cast the result into an error
// and then call the provided onError handler.
//
// Goes without saying but for this to be useful
// you must preface it with `defer`.
//
// YMMV - It mimics the goV2 check/handle proposal: https://bit.ly/354fRXv
func Handle(onError func(err error)) {
	if r := recover(); r != nil {
		e, ok := r.(error)
		if !ok {
			e = fmt.Errorf("%v", r)
		}
		onError(e)
	}
}

// Unwrap returns the result of calling the Unwrap method on err, if err's
// type contains an Unwrap method returning error.
// Otherwise, Unwrap returns nil.
//
// This is just a wrapper for the https://golang.org/pkg/errors/#Unwrap
// Provided here simply for convenience so you don't have to import
// multiple error packages.
func Unwrap(err error) error {
	return errors.Unwrap(err)
}

// Cause will unwrap the entire error chain until the root error is found.
// ie: the cause.
//
// In the case of a *goerr.Error then the cause is considered to be an instance
// that has it's callers set.
//
// If no *goerr.Error is found in the chain then this will unwrap all errors
// regardless of type.
func Cause(err error) error {
	var g *Error
	if As(err, &g) {
		for {
			unWrapped := Unwrap(g)
			if unWrapped == nil {
				break
			}
			unWrappedCasted, ok := unWrapped.(*Error)
			if !ok {
				break
			}
			if unWrappedCasted.callers == nil {
				break
			}
			g = unWrappedCasted
		}
		return g
	}

	e := err
	for {
		unWrapped := Unwrap(e)
		if unWrapped == nil {
			break
		}
		e = unWrapped
	}
	return e
}

// Is reports whether any error in err's chain matches target.
//
// The chain consists of err itself followed by the sequence of errors obtained by
// repeatedly calling Unwrap.
//
// An error is considered to match a target if it is equal to that target or if
// it implements a method Is(error) bool such that Is(target) returns true.
//
// An error type might provide an Is method so it can be treated as equivalent
// to an existing error. For example, if MyError defines
//
//	func (m MyError) Is(target error) bool { return target == os.ErrExist }
//
// then Is(MyError{}, os.ErrExist) returns true. See syscall.Errno.Is for
// an example in the standard library.
//
// This is just a wrapper for the https://golang.org/pkg/errors/#Is
// Provided here simply for convenience so you don't have to import
// multiple error packages.
func Is(err error, target error) bool {
	return errors.Is(err, target)
}

// As finds the first error in err's chain that matches target, and if so, sets
// target to that error value and returns true. Otherwise, it returns false.
//
// The chain consists of err itself followed by the sequence of errors obtained by
// repeatedly calling Unwrap.
//
// An error matches target if the error's concrete value is assignable to the value
// pointed to by target, or if the error has a method As(interface{}) bool such that
// As(target) returns true. In the latter case, the As method is responsible for
// setting target.
//
// An error type might provide an As method so it can be treated as if it were a
// different error type.
//
// As panics if target is not a non-nil pointer to either a type that implements
// error, or to any interface type.
//
// This is just a wrapper for the https://golang.org/pkg/errors/#As
// Provided here simply for convenience so you don't have to import
// multiple error packages.
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}
