# Simple

This is the hello world example, that shows off the basic functionality of this package.
Essentially use `goerr.Wrap(err)` everywhere you would normally return an error.

## Expected Output

```
crash3 received 1234567810: expecting 123456789

main.crash3:C:/Users/brad.jones/Projects/Personal/goerr/examples/simple/main.go:32
        return goerr.Wrap(errFoo, "crash3 received "+abc)
main.crash2:C:/Users/brad.jones/Projects/Personal/goerr/examples/simple/main.go:22
        return goerr.Wrap(err)
main.crash1:C:/Users/brad.jones/Projects/Personal/goerr/examples/simple/main.go:15
        return goerr.Wrap(err)
```
