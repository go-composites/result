package Result

import (
	Error "github.com/go-composites/error/src"
	NullError "github.com/go-composites/error/src/null"
	Null "github.com/go-composites/null/src"
)

type Interface interface {
	Payload() interface{}
	HasError() bool
	Error() Error.Interface
}

type data struct {
	payload interface{}
	error   Error.Interface
}

type Option func(*data)

/*
Create a new result.
*/
// okResult is the shared empty success Result (null payload, no error). A
// Result is immutable once built, so the very common no-argument completion
// case — Result.New() from Each/Clear/… — returns this one cached instance
// instead of allocating. With null/null-error interned too, a Result.New() with
// options now allocates only the wrapper itself, not its defaults.
var okResult = &data{
	payload: Null.New(),
	error:   NullError.New(),
}

func New(options ...Option) Interface {
	if len(options) == 0 {
		return okResult
	}
	d := &data{
		payload: Null.New(),
		error:   NullError.New(),
	}
	for _, opt := range options {
		opt(d)
	}
	return d
}

/*
Functional parameter to set the result payload.
*/
func WithPayload(payload interface{}) Option {
	return func(d *data) {
		d.payload = payload
	}
}

/*
Functional parameter to set the result error.
*/
func WithError(error Error.Interface) Option {
	return func(d *data) {
		d.error = error
	}
}

/*
Return the payload of an result.
*/
func (d data) Payload() interface{} {
	return d.payload
}

/*
Say if a result has an error attached.
*/
func (d data) HasError() bool {
	return !d.error.IsNull()
}

/*
Return the error of a result.
*/
func (d data) Error() Error.Interface {
	return d.error
}
