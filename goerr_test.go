package goerr_test

import (
	"fmt"
	"testing"

	"github.com/brad-jones/goerr/v2"
	"github.com/stretchr/testify/assert"
)

func TestUnwrap(t *testing.T) {
	e1 := fmt.Errorf("abc")
	e2 := fmt.Errorf("%w", e1)
	assert.Equal(t, e1, goerr.Unwrap(e2))
}

func TestCause(t *testing.T) {
	e1 := fmt.Errorf("abc")
	e2 := fmt.Errorf("%w", e1)
	e3 := fmt.Errorf("%w", e2)
	assert.Equal(t, e1, goerr.Cause(e3))
}

func TestIs(t *testing.T) {
	e1 := goerr.New("abc")
	e2 := goerr.Wrap(e1)
	e3 := goerr.Wrap(e2)
	assert.Equal(t, true, goerr.Is(e3, e1))
}

func TestAs(t *testing.T) {
	e1 := goerr.New("abc")
	e2 := goerr.Wrap(e1)
	e3 := fmt.Errorf("%w", e2)
	var err *goerr.Error
	result := goerr.As(e3, &err)
	assert.Equal(t, true, result)
	assert.Equal(t, e2, err)
}
