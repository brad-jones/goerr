# goerr

<div style="text-align:center">

[![PkgGoDev](https://pkg.go.dev/badge/github.com/brad-jones/goerr/v2)](https://pkg.go.dev/github.com/brad-jones/goerr/v2)
[![GoReport](https://goreportcard.com/badge/github.com/brad-jones/goerr/v2)](https://goreportcard.com/report/github.com/brad-jones/goerr/v2)
[![GoLang](https://img.shields.io/badge/golang-%3E%3D%201.15.1-lightblue.svg)](https://golang.org)

![.github/workflows/main.yml](https://github.com/brad-jones/goerr/workflows/.github/workflows/main.yml/badge.svg?branch=v2)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg)](https://conventionalcommits.org)
[![KeepAChangelog](https://img.shields.io/badge/Keep%20A%20Changelog-1.0.0-%23E05735)](https://keepachangelog.com/)
[![License](https://img.shields.io/github/license/brad-jones/goerr.svg)](https://github.com/brad-jones/goerr/blob/v2/LICENSE)

</div>

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
		// This will both store stackframe information and wrap the error
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

Running the above will output something similar to:

```
crash3 received 1234567810: expecting 123456789

main.crash3:C:/Users/brad.jones/Projects/Personal/goerr/examples/simple/main.go:32
	return goerr.Trace(errFoo, "crash3 received "+abc)
main.crash2:C:/Users/brad.jones/Projects/Personal/goerr/examples/simple/main.go:21
	if err := crash3(abc + "7810"); err != nil {
main.crash1:C:/Users/brad.jones/Projects/Personal/goerr/examples/simple/main.go:12
	if err := crash2(abc + "456"); err != nil {
main.main:C:/Users/brad.jones/Projects/Personal/goerr/examples/simple/main.go:38
	if err := crash1("123"); err != nil {
runtime.main:C:/Users/brad.jones/scoop/apps/go/current/src/runtime/proc.go:204
	fn()
runtime.goexit:C:/Users/brad.jones/scoop/apps/go/current/src/runtime/asm_amd64.s:1374
	BYTE    $0x90   // NOP
```

_Also see further working examples under: <https://github.com/brad-jones/goerr/tree/v2/examples>_
