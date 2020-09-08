package goerr

import (
	"encoding/json"
)

// StackTrace is an object that represents a stack trace for a given error,
// create new instances with NewStackTrace.
type StackTrace struct {
	Error    error
	Cause    error
	ErrorMsg string
	ErrorCtx map[string]interface{}
	Stack    []*StackFrame
}

// NewStackTrace is the constructor for StackTrace
func NewStackTrace(err error) *StackTrace {
	st := &StackTrace{
		Error:    err,
		Cause:    Cause(err),
		ErrorMsg: err.Error(),
	}

	if cause, ok := st.Cause.(*Error); ok {
		st.Stack = cause.Frames()
		st.ErrorCtx = marshalError(cause.Unwrap())
	}

	return st
}

// String implements the Stringer interface
func (s *StackTrace) String() string {
	st := s.ErrorMsg + "\n\n"

	if s.ErrorCtx != nil {
		ctx, err := json.MarshalIndent(s.ErrorCtx, "", "    ")
		if err != nil {
			panic(err)
		}
		st = st + string(ctx) + "\n\n"
	}

	if s.Stack != nil {
		for _, f := range s.Stack {
			st = st + f.String()
		}
		st = st + "\n"
	}

	return st
}

// MarshalJSON implements the Marshaler interface
// see https://golang.org/pkg/encoding/json/#Marshaler
func (s *StackTrace) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{
		"error-msg": s.ErrorMsg,
	}

	if s.ErrorCtx != nil {
		data["error-ctx"] = s.ErrorCtx
	}

	if s.Stack != nil {
		data["stack"] = s.Stack
	}

	return json.Marshal(data)
}

func marshalError(err error) map[string]interface{} {
	if j, jerr := json.Marshal(err); jerr == nil {
		jS := string(j)
		if jS != "" && jS != "{}" && jS != "[]" && jS != "null" {
			var out map[string]interface{}
			err := json.Unmarshal(j, &out)
			if err != nil {
				panic(err)
			}
			return out
		}
	}

	return nil
}
