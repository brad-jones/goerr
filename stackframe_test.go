package goerr_test

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/brad-jones/goerr/v2"
	"github.com/stretchr/testify/assert"
)

func TestStackFrameNew(t *testing.T) {
	cwd, err := os.Getwd()
	if assert.NoError(t, err) {
		pc, _, _, ok := runtime.Caller(0)
		if assert.Equal(t, true, ok) {
			frame := goerr.NewStackFrame(pc)
			assert.Equal(t, "github.com/brad-jones/goerr/v2_test", frame.Package)
			assert.Equal(t, "TestStackFrameNew", frame.Name)
			assert.Equal(t, filepath.Join(cwd, "stackframe_test.go"), filepath.Clean(frame.File))
			assert.Equal(t, 16, frame.LineNumber)
			assert.Equal(t, pc, frame.ProgramCounter)
		}
	}
}

func TestStackFrameSourceLine(t *testing.T) {
	pc, _, _, ok := runtime.Caller(0)
	if assert.Equal(t, true, ok) {
		frame := goerr.NewStackFrame(pc)
		line, err := frame.SourceLine()
		if assert.NoError(t, err) {
			assert.Equal(t, "pc, _, _, ok := runtime.Caller(0)", line)
		}
	}
}
