package main

import (
	"github.com/brad-jones/goerr/v2"
)

// Complex custom error types that carry additional context can of course
// also be easily defined. The following type has nothing to do with goerr
// and is just plain old go.
type httpError struct {
	StatusCode int
}

func (h *httpError) Error() string {
	return "http request failed"
}

func crash1() error {
	return goerr.Wrap(&httpError{StatusCode: 500})
}

func main() {
	if err := crash1(); err != nil {
		// You should notice you now see 3 pieces of information.
		// The error message, the status code & the full stacktrace.
		goerr.PrintTrace(err)
	}
}
