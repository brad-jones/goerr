package goerr_test

import (
	"fmt"
	"os"

	"github.com/brad-jones/goerr"
	"github.com/go-errors/errors"
)

func ExampleCheck() {
	_, err := os.Open("")
	goerr.Check(err) // expect this to panic
}

func ExampleHandle() {
	defer goerr.Handle(func(err error) {
		fmt.Println("an error")
	})
	_, err := os.Open("")
	goerr.Check(err) // expect this to panic
	// Output: an error
}

func ExampleHandleAndLog() {
	defer goerr.HandleAndLog(func(err error) {
		fmt.Println("an error")
	})
	_, err := os.Open("")
	goerr.Check(err) // expect this to panic
	// Output: open : no such file or directory
	// an error
}

func ExampleHandleAndLogWithTrace() {
	defer goerr.HandleAndLogWithTrace(func(err error) {
		fmt.Println("an error")
	})
	_, err := os.Open("")
	goerr.Check(errors.New(err)) // expect this to panic

	// open : no such file or directory
	// *os.PathError open : no such file or directory
	// /home/brad/Projects/Personal/goerr/examples_test.go:42 (0x4f2aea)
	// 	ExampleHandleAndLogWithTrace: goerr.Check(errors.New(err)) // expect this to panic
	// /home/brad/.goenv/versions/1.12.5/src/testing/example.go:121 (0x4b15ed)
	// 	runExample: eg.F()
	// /home/brad/.goenv/versions/1.12.5/src/testing/example.go:45 (0x4b1218)
	// 	runExamples: if !runExample(eg) {
	// /home/brad/.goenv/versions/1.12.5/src/testing/testing.go:1073 (0x4b4f6f)
	// 	(*M).Run: exampleRan, exampleOk := runExamples(m.deps.MatchString, m.examples)
	// _testmain.go:48 (0x4f2eee)
	// /home/brad/.goenv/versions/1.12.5/src/runtime/proc.go:200 (0x42ca6c)
	// 	main: fn()
	// /home/brad/.goenv/versions/1.12.5/src/runtime/asm_amd64.s:1337 (0x457881)
	// 	goexit: BYTE	$0x90	// NOP
	//
	// an error
}

func ExampleUnwrap() {
	originalError := &goerr.ErrUnwrappingNotSupported{}
	wrappedError := errors.New(originalError)
	unWrappedError, _ := goerr.Unwrap(wrappedError)
	fmt.Println(unWrappedError == originalError)
	// Output: true
}

func ExampleTrace() {
	st, _ := goerr.Trace(errors.New("an error"))
	fmt.Println(st)

	// *errors.errorString an error
	// /home/brad/Projects/Personal/goerr/examples_test.go:70 (0x4f2a9d)
	// 	ExampleTrace: st, _ := goerr.Trace(errors.New("an error"))
	// /home/brad/.goenv/versions/1.12.5/src/testing/example.go:121 (0x4b15ed)
	// 	runExample: eg.F()
	// /home/brad/.goenv/versions/1.12.5/src/testing/example.go:45 (0x4b1218)
	// 	runExamples: if !runExample(eg) {
	// /home/brad/.goenv/versions/1.12.5/src/testing/testing.go:1073 (0x4b4f6f)
	// 	(*M).Run: exampleRan, exampleOk := runExamples(m.deps.MatchString, m.examples)
	// _testmain.go:50 (0x4f2e7e)
	// /home/brad/.goenv/versions/1.12.5/src/runtime/proc.go:200 (0x42ca6c)
	// 	main: fn()
	// /home/brad/.goenv/versions/1.12.5/src/runtime/asm_amd64.s:1337 (0x457881)
	// 	goexit: BYTE	$0x90	// NOP
}
