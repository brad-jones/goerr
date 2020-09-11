package goerr

import (
	"encoding/json"
	"strings"
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

	// Assign any additional context values
	st.ErrorCtx = marshalError(st.Cause)

	// Grab all the frames from each error in the error chain
	frames := []*StackFrame{}
	var e *Error
	if As(err, &e) {
		for {
			frames = append(frames, e.Frame())
			unWrapped := Unwrap(e)
			if unWrapped == nil {
				break
			}
			unWrappedCasted, ok := unWrapped.(*Error)
			if !ok {
				break
			}
			if unWrappedCasted.caller == 0 {
				break
			}
			e = unWrappedCasted
		}
	}

	if len(frames) > 0 {
		// Reverse the frames so we create a call stack in the expected manner.
		// ie: from the error cause to the root of the program
		// https://github.com/golang/go/wiki/SliceTricks#reversing
		for i := len(frames)/2 - 1; i >= 0; i-- {
			opp := len(frames) - 1 - i
			frames[i], frames[opp] = frames[opp], frames[i]
		}
		st.Stack = frames
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
		if strings.HasPrefix(jS, "{") && jS != "{}" {
			var out map[string]interface{}
			if err := json.Unmarshal(j, &out); err != nil {
				return nil
			}
			return out
		}
	}
	return nil
}
