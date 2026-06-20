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
func New(options ...Option) Interface {
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
