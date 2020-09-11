package main_test

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	out, err := exec.Command("go", "run", ".").CombinedOutput()
	if assert.NoError(t, err) {
		assert.Equal(t,
			[]string{
				"crash3 received 1234567810: expecting 123456789",
				"",
				"main.crash3:/main.go:32",
				"\treturn goerr.Trace(errFoo, \"crash3 received \"+abc)",
				"main.crash2:/main.go:21",
				"\tif err := crash3(abc + \"7810\"); err != nil {",
				"main.crash1:/main.go:12",
				"\tif err := crash2(abc + \"456\"); err != nil {",
				"main.main:/main.go:38",
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
