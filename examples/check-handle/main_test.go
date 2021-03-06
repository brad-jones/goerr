package main_test

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckHandle(t *testing.T) {
	out, err := exec.Command("go", "run", ".").CombinedOutput()
	if assert.NoError(t, err) {
		assert.Equal(t,
			[]string{
				"we couldn't open the file: open /tmp/not-found/a9e5b8c7-13f6-4acc-a0c8-978319cb738b: no such file or directory",
				"",
				"main.crash1:/main.go:18",
				"\tgoerr.Check(err, \"we couldn't open the file\")",
				"main.main:/main.go:10",
				"\tif err := crash1(); err != nil {",
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
