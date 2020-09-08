package goerr_test

import (
	"fmt"
	"testing"

	"github.com/brad-jones/goerr/v2"
	"github.com/stretchr/testify/assert"
)

func TestErrorNewFromString(t *testing.T) {
	e := goerr.New("abc")
	assert.Equal(t, "abc", e.Error())
}

func TestErrorNewFromNumber(t *testing.T) {
	e := goerr.New(123)
	assert.Equal(t, "123", e.Error())
}

type Foo struct {
	Bar string
}

func TestErrorNewFromStruct(t *testing.T) {
	e := goerr.New(&Foo{Bar: "abc"})
	assert.Equal(t, "&{Bar:abc}", e.Error())
}

func TestErrorNewFromError(t *testing.T) {
	e := goerr.New(fmt.Errorf("abc"))
	assert.Equal(t, "abc", e.Error())
}

func TestErrorNewFromGoErr(t *testing.T) {
	e := goerr.New(goerr.New("abc"))
	assert.Equal(t, "abc", e.Error())
}

func TestErrorUnwrap(t *testing.T) {
	innerErr := fmt.Errorf("abc")
	e := goerr.New(innerErr)
	assert.Equal(t, innerErr, e.Unwrap())
}
