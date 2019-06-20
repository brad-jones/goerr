# goerr

[![GoReport](https://goreportcard.com/badge/brad-jones/goerr)](https://goreportcard.com/report/brad-jones/goerr)
[![GoLang](https://img.shields.io/badge/golang-%3E%3D%201.12.6-lightblue.svg)](https://golang.org)
[![GoDoc](https://godoc.org/github.com/brad-jones/goerr?status.svg)](https://godoc.org/github.com/brad-jones/goerr)
[![License](https://img.shields.io/github/license/brad-jones/goerr.svg)](https://github.com/brad-jones/goerr/blob/master/LICENSE)

This package attempts to bring proposed error handling in Go v2 into Go v1.
see: <https://go.googlesource.com/proposal/+/master/design/go2draft.md>

## Usage

`go get -u github.com/brad-jones/goerr`

```go
package copy

import (
    "os"
    "fmt"
    "io"

    . "github.com/brad-jones/goerr"
    "github.com/go-errors/errors"
)


func File(src, dst string) (err error) {
    defer Handle(func(e error){
        err = errors.Errorf("copy %s %s: %v", src, dst, e)
    })

    r, err := os.Open(src); Check(err)
    defer r.Close()

    w, err := os.Create(dst); Check(err)
    defer Handle(func(e error){
        w.Close()
        os.Remove(dst)
        panic(e)
    })

    _, err = io.Copy(w, r); Check(err)
    Check(w.Close())

    return nil
}
```