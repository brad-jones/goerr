/*
Package goerr adds additional error handling capabilities to go.

Stacktrace support

The problem with standard go errors is that often an error will bubble all the
way to the root of a program and then finally be output as a few lines of text
without any context to help diagnose the problem.

The classic case:

	open /tmp/foo/bar.xyz: The system cannot find the file specified.

You see that single line in a log file or your terminal window and unless have
an innate understanding of your program it's difficult to determin exactly where
the error was encountered.

This package borrows (literally in some cases) from other similar packages:

https://github.com/go-errors/errors

https://github.com/pkg/errors

Dave Cheney has a lot to say on the matter:

https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully

The difference is that this package has been built on top of the now standard
error wrapping support that was introduced in Go v1.13.

We can use Trace like this:

	package main

	import (
		"github.com/brad-jones/goerr/v2"
	)

	// Simple re-useable error types can be defined like this.
	// This is essentially the same as "errors.New()" but creates a `*goerr.Error`.
	var errFoo = goerr.New("expecting 123456789")

	func crash1(abc string) error {
		if err := crash2(abc + "456"); err != nil {
			// Use Trace anywhere you would normally return an error
			// This will both store stackframe infomation and wrap the error
			return goerr.Trace(err)
		}
		return nil
	}

	func crash2(abc string) error {
		if err := crash3(abc + "7810"); err != nil {
			return goerr.Trace(err)
		}
		return nil
	}

	func crash3(abc string) error {
		if abc != "123456789" {
			// Additional context messages can be added to the trace.
			// These messages should be human friendly and when prefixed
			// to the existing error message should read like a sentence.
			return goerr.Trace(errFoo, "crash3 received "+abc)
		}
		return nil
	}

	func main() {
		if err := crash1("123"); err != nil {
			goerr.PrintTrace(err)
		}
	}

And see output similar to:

	crash3 received 1234567810: expecting 123456789

	main.crash3:C:/Users/brad.jones/Projects/Personal/goerr/examples/simple/main.go:32
			return goerr.Trace(errFoo, "crash3 received "+abc)
	main.crash2:C:/Users/brad.jones/Projects/Personal/goerr/examples/simple/main.go:21
			if err := crash3(abc + "7810"); err != nil {
	main.crash1:C:/Users/brad.jones/Projects/Personal/goerr/examples/simple/main.go:12
			if err := crash2(abc + "456"); err != nil {
	main.main:C:/Users/brad.jones/Projects/Personal/goerr/examples/simple/main.go:38
			if err := crash1("123"); err != nil {
	runtime.main:C:/Users/brad.jones/scoop/apps/go/current/src/runtime/proc.go:204
			fn()
	runtime.goexit:C:/Users/brad.jones/scoop/apps/go/current/src/runtime/asm_amd64.s:1374
			BYTE    $0x90   // NOP

Check and Handle

This is totally an experiment, YMMV :)

see: https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling-overview.md#draft-design

We can emulate the proposed `check` and `handle` using `panic, defer & recover`.
Take the same example from the proposal that has been refactored to use goerr.

	import . "github.com/brad-jones/goerr/v2"

	func CopyFile(src, dst string) (err error) {
		defer Handle(func(e error){
			err = Trace(e, fmt.Sprintf("failed to copy %s to %s", src, dst))
		})

		r, err := os.Open(src); Check(err)
		defer r.Close()

		w, err := os.Create(dst); Check(err)
		defer Handle(func(e error){
			w.Close()
			os.Remove(dst)
			panic(e) // re-panic to make above handler set the err
		})

		_, err = io.Copy(w, r); Check(err)
		Check(w.Close())

		return nil
	}

So `Check()` replaces the repetitive `if err != nil { ... }` phrase and
`Handle` takes care of the `recover()` logic for you. `Check()` automatically
calls `Trace()` on your error.

Yeah I get it this looks like exceptions and if you choose to use it like that
then thats your prerogative, I'm not going to stop you... but you probably
shouldn't!

Panicing doesn't work across goroutines for a start and this
https://go101.org/article/panic-and-recover-more.html

I think where this can be really useful is when you say have a function like this:

	func DoSomeWork() (err error) {
		defer Handle(func(e error){
			err = e
		})
		Check(build("foo"))
		Check(build("bar"))
		Check(build("baz"))
		Check(build("qux"))
	}
*/
package goerr
