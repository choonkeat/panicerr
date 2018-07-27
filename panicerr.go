package panicerr

import (
	"fmt"
	"runtime"

	"github.com/pkg/errors"
)

type panicErr struct {
	r   interface{}
	pcs []uintptr
}

// New returns a pkg/errors compatible `errors.WithStack` value
// ref code from https://github.com/pkg/errors/blob/816c9085562cd7ee03e7f8188a1cfd942858cded/stack.go#L125-L139
func New(recovered interface{}) error {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(5, pcs[:])
	return panicErr{r: recovered, pcs: pcs[0:n]}
}

// Error implements `error`
func (err panicErr) Error() string {
	return fmt.Sprintf("%s", err.r)
}

// StackTrace implements `errors.WithStack`
func (err panicErr) StackTrace() errors.StackTrace {
	frames := make([]errors.Frame, len(err.pcs))
	for i := range err.pcs {
		frames[i] = errors.Frame(err.pcs[i])
	}
	return frames
}
