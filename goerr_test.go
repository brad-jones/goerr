package goerr_test

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"

	"github.com/brad-jones/goerr/v2"
	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	out, err := exec.Command("go", "run", "./examples/simple").CombinedOutput()
	if assert.NoError(t, err) {
		assert.Equal(t,
			[]string{
				"crash3 received 1234567810: expecting 123456789",
				"",
				"main.crash3:/examples/simple/main.go:32",
				"\treturn goerr.Trace(errFoo, \"crash3 received \"+abc)",
				"main.crash2:/examples/simple/main.go:21",
				"\tif err := crash3(abc + \"7810\"); err != nil {",
				"main.crash1:/examples/simple/main.go:12",
				"\tif err := crash2(abc + \"456\"); err != nil {",
				"main.main:/examples/simple/main.go:38",
				"\tif err := crash1(\"123\"); err != nil {",
				"runtime.main:/src/runtime/proc.go:204",
				"\tfn()",
				"runtime.goexit:/src/runtime/asm_amd64.s:1374",
				"\tBYTE\t$0x90\t// NOP",
				"",
				"",
			},
			normaliseCmdOutput(out),
		)
	}
}

func TestCustom(t *testing.T) {
	out, err := exec.Command("go", "run", "./examples/custom").CombinedOutput()
	if assert.NoError(t, err) {
		assert.Equal(t,
			[]string{
				"http request failed",
				"",
				"{",
				"    \"StatusCode\": 500",
				"}",
				"",
				"main.crash1:/examples/custom/main.go:19",
				"\treturn goerr.Trace(&httpError{StatusCode: 500})",
				"main.main:/examples/custom/main.go:23",
				"\tif err := crash1(); err != nil {",
				"runtime.main:/src/runtime/proc.go:204",
				"\tfn()",
				"runtime.goexit:/src/runtime/asm_amd64.s:1374",
				"\tBYTE\t$0x90\t// NOP",
				"",
				"",
			},
			normaliseCmdOutput(out),
		)
	}
}

func TestCheckHandle(t *testing.T) {
	out, err := exec.Command("go", "run", "./examples/check-handle").CombinedOutput()
	if assert.NoError(t, err) {
		assert.Equal(t,
			[]string{
				"crash1 failed because: we couldn't open the file: open /tmp/not-found/a9e5b8c7-13f6-4acc-a0c8-978319cb738b: no such file or directory",
				"",
				"{",
				"    \"Err\": 2,",
				"    \"Op\": \"open\",",
				"    \"Path\": \"/tmp/not-found/a9e5b8c7-13f6-4acc-a0c8-978319cb738b\"",
				"}",
				"",
				"main.crash1:/examples/check-handle/main.go:23",
				"\tgoerr.Check(err, \"we couldn't open the file\")",
				"main.main:/examples/check-handle/main.go:13",
				"\tif err := crash1(); err != nil {",
				"runtime.main:/src/runtime/proc.go:204",
				"\tfn()",
				"runtime.goexit:/src/runtime/asm_amd64.s:1374",
				"\tBYTE\t$0x90\t// NOP",
				"",
				"",
			},
			normaliseCmdOutput(out),
		)
	}
}

func normaliseCmdOutput(in []byte) []string {
	root := strings.ReplaceAll(runtime.GOROOT(), "\\", "/")
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cwd = strings.ReplaceAll(cwd, "\\", "/")

	out := string(in)
	out = strings.ReplaceAll(out, "\r\n", "\n")
	out = strings.ReplaceAll(out, root, "")
	out = strings.ReplaceAll(out, cwd, "")
	out = strings.ReplaceAll(out, "    \"Err\": 3,", "    \"Err\": 2,")
	out = strings.ReplaceAll(out, "The system cannot find the path specified.", "no such file or directory")

	return strings.Split(out, "\n")
}

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

func TestCauseGoErr(t *testing.T) {
	e1 := goerr.New("abc")
	e2 := goerr.Trace(e1)
	e3 := goerr.Trace(e2)
	assert.Equal(t, goerr.Unwrap(e2), goerr.Cause(e3))
}

func TestIs(t *testing.T) {
	e1 := goerr.New("abc")
	e2 := goerr.Trace(e1)
	e3 := goerr.Trace(e2)
	assert.Equal(t, true, goerr.Is(e3, e1))
}

func TestAs(t *testing.T) {
	e1 := goerr.New("abc")
	e2 := goerr.Trace(e1)
	e3 := fmt.Errorf("%w", e2)
	var err *goerr.Error
	result := goerr.As(e3, &err)
	assert.Equal(t, true, result)
	assert.Equal(t, e2, err)
}
