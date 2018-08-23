# Accessor

generate getter and setter for unexported struct field

# Install

```
$ go get -u github.com/mizkei/accessor
```

# Example

```go
package types

import (
	"time"
)

// Example is example
//go:generate accessor -setter -type=Example
type Example struct {
	A int
	b string
	c time.Time
}
```

execute `go generate`

output
```go
// Code generated by "accessor -setter -type=Example"; DO NOT EDIT.

package types

import "time"

// B return b value
func (t Example) B() string {
	return t.b
}

// SetB set v to b
func (t *Example) SetB(v string) {
	t.b = v
}

// C return c value
func (t Example) C() time.Time {
	return t.c
}

// SetC set v to c
func (t *Example) SetC(v time.Time) {
	t.c = v
}
```

# TODO

- test
- generate anonymous struct type
- use field tag