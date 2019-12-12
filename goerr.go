package goerr

import (
	"errors"
	"fmt"

	errors2 "github.com/go-errors/errors"
	"github.com/go-log/log"
	errors1 "github.com/pkg/errors"
)

// Goerr is a class like object, create new instances with `goerr.New()`
type Goerr struct {
	logger log.Logger
}

// New creates new instances of `Goerr`.
//
// The logger must implement the interface from https://github.com/go-log/log
func New(logger log.Logger) *Goerr {
	return &Goerr{
		logger: logger,
	}
}

// Check will panic if err is not null
func (g *Goerr) Check(err error) {
	if err != nil {
		panic(g.Wrap(err))
	}
}

// Handle will recover, cast the result into an error
// and then call the provided onError handler.
//
// Goes without saying but for this to be useful
// you must preface it with `defer`.
func (g *Goerr) Handle(onError func(err error)) {
	if r := recover(); r != nil {
		var ok bool
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("%v", r)
		}
		onError(err)
	}
}

// HandleAndLog does the same thing as Handle but also logs (using the provided
// logger) the error message.
func (g *Goerr) HandleAndLog(onError func(err error)) {
	if r := recover(); r != nil {
		var ok bool
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("%v", r)
		}
		g.logger.Logf("%v\n", err.Error())
		onError(err)
	}
}

// HandleAndLogWithTrace does the same thing as HandleAndLog but also logs
// (using the provided logger) a stack trace that was attached to the error.
//
// If Trace returns an error nothing will be logged and it will silently fail.
// This would be the case if you are handling a non wrapped error.
func (g *Goerr) HandleAndLogWithTrace(onError func(err error)) {
	if r := recover(); r != nil {
		var ok bool
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("%v", r)
		}
		g.logger.Logf("%v\n", err.Error())
		if st, err := g.Trace(err); err == nil {
			g.logger.Logf("%v\n", st)
		}
		onError(err)
	}
}

// Errorf is just an alias to "github.com/go-errors/errors".
func (g *Goerr) Errorf(format string, a ...interface{}) error {
	return errors2.Errorf(format, a...)
}

// WrapPrefix is just an alias to "github.com/go-errors/errors".
func (g *Goerr) WrapPrefix(err error, prefix string) error {
	return errors2.WrapPrefix(err, prefix, 1)
}

// Wrap is just an alias to "github.com/go-errors/errors".
func (g *Goerr) Wrap(err error) error {
	return errors2.Wrap(err, 1)
}

// Unwrap takes an error value and assumes it is either a https://github.com/go-errors/errors
// object or a https://github.com/pkg/errors object and attempts to unwrap the error returning
// the original untouched error value.
//
// If the error can not be unwrapped, the second error returned will be a
// wrapped *ErrUnwrappingNotSupported object.
func (g *Goerr) Unwrap(err error) (error, error) {
	if err == nil {
		return nil, nil
	}
	type causer interface {
		Cause() error
	}
	unwrapStdLib := func(err error) error {
		if v := errors.Unwrap(err); v != nil {
			return v
		}
		return err
	}
	switch v := err.(type) {
	case causer:
		return unwrapStdLib(v.Cause()), nil
	case *errors2.Error:
		unWrapped := v.Err
		for {
			if v, ok := unWrapped.(*errors2.Error); ok {
				unWrapped = v
			} else {
				break
			}
		}
		return unwrapStdLib(unWrapped), nil
	}
	if e := unwrapStdLib(err); e != nil {
		return e, nil
	}
	return nil, errors2.New(&ErrUnwrappingNotSupported{
		OriginalError: err,
	})
}

// MustUnwrap does the same thing as Unwrap but panics instead of returning a second error.
func (g *Goerr) MustUnwrap(err error) error {
	v, err := g.Unwrap(err)
	g.Check(err)
	return v
}

// Trace takes an error value and assumes it is either a https://github.com/go-errors/errors
// object or a https://github.com/pkg/errors object and attempts to extract a
// stack trace from the error value, returning it as a string.
//
// If the error does not appear to have stack trace attached, this will return
// a wrapped *ErrUnwrappingNotSupported object.
func (g *Goerr) Trace(err error) (string, error) {
	type stackTracer interface {
		StackTrace() errors1.StackTrace
	}
	type errorStacker interface {
		ErrorStack() string
	}
	switch v := err.(type) {
	case stackTracer:
		return fmt.Sprintf("%+v\n", v.StackTrace()), nil
	case errorStacker:
		return v.ErrorStack(), nil
	}
	return "", errors2.New(&ErrStackTraceNotSupported{
		OriginalError: err,
	})
}

// MustTrace does the same thing as Trace but panics instead of returning an error.
func (g *Goerr) MustTrace(err error) string {
	v, err := g.Trace(err)
	g.Check(err)
	return v
}
