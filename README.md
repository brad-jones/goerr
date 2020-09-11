# goerr

[![PkgGoDev](https://pkg.go.dev/badge/github.com/brad-jones/goerr/v2)](https://pkg.go.dev/github.com/brad-jones/goerr/v2)
[![GoReport](https://goreportcard.com/badge/github.com/brad-jones/goerr/v2)](https://goreportcard.com/report/github.com/brad-jones/goerr/v2)
[![GoLang](https://img.shields.io/badge/golang-%3E%3D%201.15.1-lightblue.svg)](https://golang.org)
![.github/workflows/main.yml](https://github.com/brad-jones/goerr/workflows/.github/workflows/main.yml/badge.svg?branch=v2)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg)](https://conventionalcommits.org)
[![KeepAChangelog](https://img.shields.io/badge/Keep%20A%20Changelog-1.0.0-%23E05735)](https://keepachangelog.com/)
[![License](https://img.shields.io/github/license/brad-jones/goerr.svg)](https://github.com/brad-jones/goerr/blob/v2/LICENSE)

Package goerr adds additional error handling capabilities to go.

_Looking for v1, see the [master branch](https://github.com/brad-jones/goerr/tree/master)_

## Quick Start

`go get -u github.com/brad-jones/goerr/v2`

```go
package main

import (
	"github.com/brad-jones/goerr/v2"
)

// Simple re-useable error types can be defined like this.
// This is essentially the same as "errors.New()" but creates a `*goerr.Error`.
var errFoo = goerr.New("expecting 123456789")

func crash1(abc string) error {
	if err := crash2(abc + "456"); err != nil {
		// Use Wrap anywhere you would normally return an error
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
```

Running the above will output something similar to:

```
crash3 received 1234567810: expecting 123456789

main.crash3:C:/Users/brad.jones/Projects/Personal/goerr/examples/simple/main.go:32
    return goerr.Wrap(errFoo, "crash3 received "+abc)
main.crash2:C:/Users/brad.jones/Projects/Personal/goerr/examples/simple/main.go:22
    return goerr.Wrap(err)
main.crash1:C:/Users/brad.jones/Projects/Personal/goerr/examples/simple/main.go:15
    return goerr.Wrap(err)
```

_Also see further working examples under: <https://github.com/brad-jones/goerr/tree/v2/examples>_
