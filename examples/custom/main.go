package main

import (
	"github.com/brad-jones/goerr/v2"
)

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
		goerr.PrintTrace(err)
	}
}
