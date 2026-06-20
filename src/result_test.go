package Result

import (
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
