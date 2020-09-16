package main

import (
	"github.com/brad-jones/goerr/v2"
)

var errFoo = goerr.New("expecting 123456789")

func crash1(abc string) error {
	if err := crash2(abc + "456"); err != nil {
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
		return goerr.Wrap(errFoo, "crash3 received "+abc)
	}
	return nil
}

func main() {
	if err := crash1("123"); err != nil {
		goerr.PrintTrace(err)
	}
}
