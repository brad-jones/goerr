/*
Package goerr attempts to bring proposed error handling in Go v2 into Go v1.

see: https://go.googlesource.com/proposal/+/master/design/go2draft.md

While some of my idea's presented here may not be academically perfect I am
taking a pragmatic approach to my development. But also on the other hand I
have come to realise that golang is it's own language and not another...
yes go is verbose suck it up and move on.

Check and Handle

We can emulate the proposed `check` and `handle` using `panic, defer & recover`.
Take the same example from the proposal that has been refactored to use goerr.

	import . "github.com/brad-jones/goerr"

	func CopyFile(src, dst string) (err error) {
		defer Handle(func(e error){
			err = fmt.Errorf("copy %s %s: %v", src, dst, e)
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

So `Check()` replaces the repetitive `if err != nil { ... }` phrase and `Handle`
takes care of the `recover()` logic for you.

Idiomatic Must Functions

It is idiomatic in golang for functions that might panic to be prefixed with
`Must`. I am going to take things a step further though and suggest that life
could be much easier if every (or most) function also had a `MustFoo()`
equivalent.

Adding the extra wrapping function is minimal effort for the API developer
(5 lines) but now gives the API consumer ultimate choice in how they want to
deal with errors.

Another example that eliminates the `Check` logic:

	import . "github.com/brad-jones/goerr"

	func CopyFile(src, dst string) (err error) {
		defer Handle(func(e error){
			err = fmt.Errorf("copy %s %s: %v", src, dst, e)
		})

		r := os.MustOpen(src);
		defer r.MustClose()

		w := os.MustCreate(dst);
		defer Handle(func(e error){
			w.MustClose()
			os.MustRemove(dst)
			panic(e) // re-panic to make above handler set the err
		})

		io.MustCopy(w, r)
		w.MustClose()

		return nil
	}

Oh No Exceptions

This looks awefully like exceptions that the go designers are purposefully
avoiding. We could go one step further with our example:

	import . "github.com/brad-jones/goerr"

	func MustCopyFile(src, dst string) {
		r := os.MustOpen(src);
		defer r.MustClose()

		w := os.MustCreate(dst);
		defer Handle(func(e error){
			w.MustClose()
			os.MustRemove(dst)
			panic(e)
		})

		io.MustCopy(w, r)
		w.MustClose()

		return nil
	}

And this is where my pragmatic approach kicks in, firstly the above still
communicates that it might panic so the consumer should be ready for such
a possibility.

It is idiomatic go to not panic across package boundaries and for the most part
I agree with this but I am going to extend this by saying that you should not
panic across solution boundaries.

Panicking with-in a solution or application (that could be split into many
packages) I believe is fine as you know when your going to panic and when you
need to recover.

Other Thoughts re Panicking

I am sure the performance of panic, defer & recover is slower than just
checking and returning an error value. Unless I am writing some super duper
performance oriented thing I doubt I'll notice any impacts.

Panicking across goroutines, yep I get it, it does not work.
I am totally fine with this. There are packages like
https://godoc.org/golang.org/x/sync/errgroup to handle such cases.

Recover doesn't always work https://go101.org/article/panic-and-recover-more.html
This nearly made me drop this entire project but I am going to persevere and
see how this package works out.

Wrapping of Errors

The other issue that will hopefully get solved in go v2 is the ability to
provide context to errors as they get passed through the stack. For now we
have solutions such as:

https://godoc.org/github.com/pkg/errors

https://godoc.org/github.com/go-errors/errors

This package provides some helpful functions to iron some of the differences
between these error wrappers (I started using pkg/errors but now prefer
go-errors/errors).

	import "github.com/go-errors/errors"
	import . "github.com/brad-jones/goerr"

	type anError struct {
		message string
	}

	func (e *anError) Error() string {
		return e.message
	}

	func foo() error {
		return errors.New(&anError{
			message: "an error happened",
		})
	}

	func bar() error {
		return errors.WrapPrefix(foo(), "some extra context", 0)
	}

	func main() {
		defer Handle(func(e error){
			switch err := MustUnwrap(e).(type) {
			case *anError:
				fmt.Println(err.message)
			default:
				fmt.Println(MustTrace(e))
			}
		})
		Check(bar())
	}

Helper Functions vs New Instance

You can use goerr via 2 different APIs, all the examples to date have been using
the simple helper functions but all these do is call the instance methods of a
`goerr.Goerr` object.

If you need to set your own logger (and maybe other things in the future) this
is how you would do it.

	import logger "github.com/go-log/log/fmt"
	import "github.com/brad-jones/goerr"

	func main() {
		l := logger.New()
		g := goerr.New(l)
		defer g.HandleAndLog(func(e error){ })
	}

see: https://github.com/go-log/log

NOTE: We have also been using a "dot" import this is not necessarily suggested
either, if your not familiar with this read up here -
https://scene-si.org/2018/01/25/go-tips-and-tricks-almost-everything-about-imports/
*/
package goerr
