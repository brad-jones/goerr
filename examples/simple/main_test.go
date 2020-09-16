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
				"main.crash3:/main.go:25",
				"\treturn goerr.Wrap(errFoo, \"crash3 received \"+abc)",
				"main.crash2:/main.go:18",
				"\treturn goerr.Wrap(err)",
				"main.crash1:/main.go:11",
				"\treturn goerr.Wrap(err)",
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
