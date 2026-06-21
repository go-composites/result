package Result

import (
	"reflect"
	"testing"

	Error "github.com/go-composites/error/src"
)

func TestNewDefault(t *testing.T) {
	r := New()
	if r == nil {
		t.Fatal("New() returned nil")
	}
	if r.HasError() {
		t.Errorf("HasError() = true, want false for a default result")
	}
	if r.Payload() == nil {
		t.Errorf("Payload() = nil, want a non-nil null value")
	}
	if r.Error() == nil {
		t.Errorf("Error() = nil, want a non-nil null error")
	}
	if !r.Error().IsNull() {
		t.Errorf("Error().IsNull() = false, want true for a default result")
	}
}

func TestWithPayload(t *testing.T) {
	want := "payload-value"
	r := New(WithPayload(want))
	if got := r.Payload(); got != want {
		t.Errorf("Payload() = %v, want %v", got, want)
	}
	if r.HasError() {
		t.Errorf("HasError() = true, want false when only a payload is set")
	}
}

func TestWithError(t *testing.T) {
	want := "boom"
	e := Error.New(want)
	r := New(WithError(e))
	if !r.HasError() {
		t.Errorf("HasError() = false, want true when an error is set")
	}
	if got := r.Error(); got != e {
		t.Errorf("Error() = %v, want %v", got, e)
	}
	if got := r.Error().Message(); got != want {
		t.Errorf("Error().Message() = %q, want %q", got, want)
	}
}

func TestOkBoolInterned(t *testing.T) {
	rTrue := Ok(true)
	if got := rTrue.Payload(); got != true {
		t.Errorf("Ok(true).Payload() = %v, want true", got)
	}
	if rTrue.HasError() {
		t.Errorf("Ok(true).HasError() = true, want false")
	}
	if Ok(true) != rTrue {
		t.Errorf("Ok(true) is not interned: repeated calls returned distinct values")
	}

	rFalse := Ok(false)
	if got := rFalse.Payload(); got != false {
		t.Errorf("Ok(false).Payload() = %v, want false", got)
	}
	if rFalse.HasError() {
		t.Errorf("Ok(false).HasError() = true, want false")
	}
	if Ok(false) != rFalse {
		t.Errorf("Ok(false) is not interned: repeated calls returned distinct values")
	}

	if rTrue == rFalse {
		t.Errorf("Ok(true) and Ok(false) returned the same instance")
	}
}

func TestOkNonBool(t *testing.T) {
	want := "payload-value"
	r := Ok(want)
	if got := r.Payload(); got != want {
		t.Errorf("Ok(%q).Payload() = %v, want %v", want, got, want)
	}
	if r.HasError() {
		t.Errorf("Ok(%q).HasError() = true, want false", want)
	}
	if !r.Error().IsNull() {
		t.Errorf("Ok(%q).Error().IsNull() = false, want true", want)
	}
}

func TestErr(t *testing.T) {
	e := Error.New("boom")
	r := Err(e)
	if !r.HasError() {
		t.Errorf("Err(e).HasError() = false, want true")
	}
	if got := r.Error(); got != e {
		t.Errorf("Err(e).Error() = %v, want %v", got, e)
	}
	// The payload is a null value: same concrete type as the one New(WithError)
	// leaves in place, and never nil.
	if r.Payload() == nil {
		t.Errorf("Err(e).Payload() = nil, want a null value")
	}
	reference := New(WithError(e))
	if got, want := reflect.TypeOf(r.Payload()), reflect.TypeOf(reference.Payload()); got != want {
		t.Errorf("Err(e).Payload() type = %v, want %v (null)", got, want)
	}
}

func TestWithPayloadAndError(t *testing.T) {
	payload := 42
	e := Error.New("failure")
	r := New(WithPayload(payload), WithError(e))
	if got := r.Payload(); got != payload {
		t.Errorf("Payload() = %v, want %v", got, payload)
	}
	if !r.HasError() {
		t.Errorf("HasError() = false, want true")
	}
	if got := r.Error(); got != e {
		t.Errorf("Error() = %v, want %v", got, e)
	}
}
