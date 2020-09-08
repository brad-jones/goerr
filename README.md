# goerr

[![GoReport](https://goreportcard.com/badge/github.com/brad-jones/goerr/v2)](https://goreportcard.com/report/github.com/brad-jones/goerr/v2)
[![GoLang](https://img.shields.io/badge/golang-%3E%3D%201.15.1-lightblue.svg)](https://golang.org)
[![GoDoc](https://godoc.org/github.com/brad-jones/goerr/v2?status.svg)](https://godoc.org/github.com/brad-jones/goerr/v2)
[![License](https://img.shields.io/github/license/brad-jones/goerr.svg)](https://github.com/brad-jones/goerr/blob/v2/LICENSE)

Package goerr adds additional error handling capabilities to go.

## Quick Start

`go get -u github.com/brad-jones/goerr/v2`

```go
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
		// This will both store stackframe infomation and wrap the error
		return goerr.Trace(err)
	}
	return nil
}

func crash2(abc string) error {
	if err := crash3(abc + "7810"); err != nil {
		return goerr.Trace(err)
	}
	return nil
}

func crash3(abc string) error {
	if abc != "123456789" {
		// Additional context messages can be added to the trace.
		// These messages should be human friendly and when prefixed
		// to the existing error message should read like a sentence.
		return goerr.Trace(errFoo, "crash3 received "+abc)
	}
	return nil
}

func main() {
	if err := crash1("123"); err != nil {
		goerr.PrintTrace(err)
	}
}

```

## Documentation

<https://pkg.go.dev/github.com/brad-jones/goerr/v2>

Also see further working examples under:
<https://github.com/brad-jones/goerr/v2/tree/master/examples>
