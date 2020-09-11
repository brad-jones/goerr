package main_test

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustom(t *testing.T) {
	out, err := exec.Command("go", "run", ".").CombinedOutput()
	if assert.NoError(t, err) {
		assert.Equal(t,
			[]string{
				"http request failed",
				"",
				"{",
				"    \"StatusCode\": 500",
				"}",
				"",
				"main.crash1:/main.go:19",
				"\treturn goerr.Trace(&httpError{StatusCode: 500})",
				"main.main:/main.go:23",
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

	return strings.Split(out, "\n")
}
