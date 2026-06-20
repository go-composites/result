<p align="center"><img src="https://raw.githubusercontent.com/go-composites/brand/main/social/go-composites.png" alt="go-composites/result" width="720"></p>

# result

[![ci](https://github.com/go-composites/result/actions/workflows/ci.yml/badge.svg)](https://github.com/go-composites/result/actions/workflows/ci.yml)

The result composite of [go-composites](https://github.com/go-composites) — the
value every fallible operation returns. A `Result` pairs a payload with an
[`Error`](https://github.com/go-composites/error), so success and failure travel
together as one value: no `(value, ok)`, no panics, no bare `nil`. A fresh
`Result.New()` defaults to a [`Null`](https://github.com/go-composites/null)
payload and a `NullError`, so it represents a successful, empty outcome.

## Install

```sh
go get github.com/go-composites/result
```

## API

| symbol | returns | notes |
| --- | --- | --- |
| `New(opts...)` | `Result.Interface` | payload defaults to `Null`, error to `NullError` |
| `WithPayload(v)` | `Option` | functional option setting the payload |
| `WithError(e)` | `Option` | functional option setting the `Error.Interface` |
| `Payload()` | `interface{}` | the carried payload |
| `HasError()` | `bool` | `!Error.IsNull()` — `true` only when a real error is attached |
| `Error()` | `Error.Interface` | the carried error (a `NullError` when there is none) |

## Usage

```go
package main

import (
	"fmt"

	Result "github.com/go-composites/result/src"
	Error "github.com/go-composites/error/src"
)

func main() {
	ok := Result.New(Result.WithPayload(42))
	fmt.Println(ok.HasError(), ok.Payload()) // false 42

	bad := Result.New(Result.WithError(Error.New("key not found")))
	fmt.Println(bad.HasError(), bad.Error().Message()) // true key not found

	empty := Result.New()
	fmt.Println(empty.HasError()) // false  (Null payload + NullError)
}
```

## License

BSD-3-Clause © the go-composites/result authors.
