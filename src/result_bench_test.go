package Result

import "testing"

// someInterfaceValue is a non-bool payload boxed in an interface, so both the
// New(WithPayload(...)) and Ok(...) paths carry an identical interface value.
var someInterfaceValue interface{} = "payload-value"

// sink keeps the constructed Result from being optimised away.
var sink Interface

func BenchmarkNewWithPayload(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sink = New(WithPayload(someInterfaceValue))
	}
}

func BenchmarkOk(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sink = Ok(someInterfaceValue)
	}
}

func BenchmarkOkBool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sink = Ok(true)
	}
}
