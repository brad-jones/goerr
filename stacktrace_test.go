package goerr_test

import (
	"fmt"
	"testing"

	"github.com/brad-jones/goerr/v2"
	"github.com/stretchr/testify/assert"
)

func TestStackTraceNewFromError(t *testing.T) {
	e := fmt.Errorf("abc")
	st := goerr.NewStackTrace(e)
	assert.Equal(t, e, st.Error)
	assert.Equal(t, e, st.Cause)
	assert.Equal(t, "abc", st.ErrorMsg)
	assert.Nil(t, st.ErrorCtx)
	assert.Nil(t, st.Stack)
}

func TestStackTraceNewFromGoErr(t *testing.T) {
	err := goerr.New("abc")
	traced := goerr.Wrap(err)
	st := goerr.NewStackTrace(traced)
	assert.Equal(t, traced, st.Error)
	assert.Equal(t, goerr.Unwrap(traced), st.Cause)
	assert.Equal(t, "abc", st.ErrorMsg)
	assert.Equal(t, 3, len(st.Stack))
	assert.Nil(t, st.ErrorCtx)
}

type fooError struct {
	Bar string
}

func (f *fooError) Error() string {
	return "a message"
}

func TestStackTraceWithCtx(t *testing.T) {
	err := goerr.Wrap(&fooError{Bar: "abc"})
	st := goerr.NewStackTrace(err)
	assert.Equal(t, map[string]interface{}{
		"Bar": "abc",
	}, st.ErrorCtx)
}
