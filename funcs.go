package goerr

import (
	"fmt"

	logger "github.com/go-log/log/fmt"
)

var defaultLogger = logger.New()
var defaultInstance = New(defaultLogger)

// Check uses the `defaultInstance` with a preconfigured logger.
// It does the same thing as `goerr.New(logger).Check()`.
func Check(err error) {
	if err != nil {
		panic(Wrap(err))
	}
}

// Handle uses the `defaultInstance` with a preconfigured logger.
// It does the same thing as `goerr.New(logger).Handle()`.
func Handle(onError func(err error)) {
	if r := recover(); r != nil {
		var ok bool
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("%v", r)
		}
		onError(err)
	}
}

// HandleAndLog uses the `defaultInstance` with a preconfigured logger.
// It does the same thing as `goerr.New(logger).HandleAndLog()`.
func HandleAndLog(onError func(err error)) {
	if r := recover(); r != nil {
		var ok bool
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("%v", r)
		}
		defaultLogger.Logf("%v\n", err.Error())
		onError(err)
	}
}

// HandleAndLogWithTrace uses the `defaultInstance` with a preconfigured logger.
// It does the same thing as `goerr.New(logger).HandleAndLogWithTrace()`.
func HandleAndLogWithTrace(onError func(err error)) {
	if r := recover(); r != nil {
		var ok bool
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("%v", r)
		}
		defaultLogger.Logf("%v\n", err.Error())
		if st, err := Trace(err); err == nil {
			defaultLogger.Logf("%v\n", st)
		} else {
			defaultLogger.Logf("%v\n", err)
		}
		onError(err)
	}
}

// Errorf uses the `defaultInstance` with a preconfigured logger.
// It does the same thing as `goerr.New(logger).Errorf()`.
func Errorf(format string, a ...interface{}) error {
	return defaultInstance.Errorf(format, a...)
}

// WrapPrefix uses the `defaultInstance` with a preconfigured logger.
// It does the same thing as `goerr.New(logger).WrapPrefix()`.
func WrapPrefix(err error, prefix string) error {
	return defaultInstance.WrapPrefix(err, prefix)
}

// Wrap uses the `defaultInstance` with a preconfigured logger.
// It does the same thing as `goerr.New(logger).Wrap()`.
func Wrap(err error) error {
	return defaultInstance.Wrap(err)
}

// Unwrap uses the `defaultInstance` with a preconfigured logger.
// It does the same thing as `goerr.New(logger).Unwrap()`.
func Unwrap(err error) (error, error) {
	return defaultInstance.Unwrap(err)
}

// MustUnwrap uses the `defaultInstance` with a preconfigured logger.
// It does the same thing as `goerr.New(logger).MustUnwrap()`.
func MustUnwrap(err error) error {
	return defaultInstance.MustUnwrap(err)
}

// Trace uses the `defaultInstance` with a preconfigured logger.
// It does the same thing as `goerr.New(logger).Trace()`.
func Trace(err error) (string, error) {
	return defaultInstance.Trace(err)
}

// MustTrace uses the `defaultInstance` with a preconfigured logger.
// It does the same thing as `goerr.New(logger).MustTrace()`.
func MustTrace(err error) string {
	return defaultInstance.MustTrace(err)
}
