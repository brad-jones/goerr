package main

import (
	"os"

	"github.com/brad-jones/goerr/v2"
)

func main() {
	if err := crash1(); err != nil {
		goerr.PrintTrace(err)
	}
}

func crash1() (err error) {
	defer goerr.Handle(func(e error) { err = e })
	f, err := os.Open("/tmp/not-found/a9e5b8c7-13f6-4acc-a0c8-978319cb738b")
	goerr.Check(err, "we couldn't open the file")
	goerr.Check(f.Close(), "we failed to close file handle")
	return
}
