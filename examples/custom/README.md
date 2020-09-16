# Custom Error Types

Complex custom error types that carry additional context can of course also be
easily defined. In go an error is any type that implements `Error() string`.

When you ask `goerr` to print a stack trace it will attempt to marshal the
error chain's  _"cause"_ into JSON, revealing even more contextual information.

You should notice in this example there are 3 pieces of information.

- The error message,
- the status code &
- the full stacktrace.

## Expected Output

```
http request failed

{
    "StatusCode": 500
}

main.crash1:C:/Users/brad.jones/Projects/Personal/goerr/examples/custom/main.go:19
        return goerr.Wrap(&httpError{StatusCode: 500})
```
