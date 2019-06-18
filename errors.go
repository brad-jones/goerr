package goerr

// ErrUnwrappingNotSupported is returned by Unwrap and panic'ed by MustUnwrap
type ErrUnwrappingNotSupported struct {
	OriginalError error
}

func (e *ErrUnwrappingNotSupported) Error() string {
	return "goerr: does not support unwrapping this error"
}

// ErrStackTraceNotSupported is returned by Trace and panic'ed by MustTrace
type ErrStackTraceNotSupported struct {
	OriginalError error
}

func (e *ErrStackTraceNotSupported) Error() string {
	return "goerr: does not support extracting a stack trace from this error"
}
