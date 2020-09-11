package main

import (
	"github.com/brad-jones/goerr/v2"
)

// Simple re-useable error types can be defined like this.
// This is essentiually the same as "errors.New()" but creates a `*goerr.Error`.
var errFoo = goerr.New("expecting 123456789")

func crash1(abc string) error {
	if err := crash2(abc + "456"); err != nil {
		// Use Trace anywhere you would normally return an error
		// This will both store stackframe information and wrap the error
		return goerr.Wrap(err)
	}
	return nil
}

func crash2(abc string) error {
	if err := crash3(abc + "7810"); err != nil {
		return goerr.Wrap(err)
	}
	return nil
}

func crash3(abc string) error {
	if abc != "123456789" {
		// Additional context messages can be added to the trace.
		// These messages should be human friendly and when prefixed
		// to the existing error message should read like a sentence.
		return goerr.Wrap(errFoo, "crash3 received "+abc)
	}
	return nil
}

func main() {
	if err := crash1("123"); err != nil {
		goerr.PrintTrace(err)
	}
}
