# Check & Handle

This library also has experimental support for the go v2 check/handle proposal
<https://bit.ly/354fRXv> through a set of functions that mimic the proposal by
using panic & recover, YMMV...

## Expected Output

```
we couldn't open the file: open /tmp/not-found/a9e5b8c7-13f6-4acc-a0c8-978319cb738b: The system cannot find the path specified.

main.crash1:C:/Users/brad.jones/Projects/Personal/goerr/examples/check-handle/main.go:18
        goerr.Check(err, "we couldn't open the file")
main.main:C:/Users/brad.jones/Projects/Personal/goerr/examples/check-handle/main.go:10
        if err := crash1(); err != nil {
```
