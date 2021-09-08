package tests

import (
  "testing"
  "fmt"
)

// AssertEqual ensures that two types are equal
// Credit to https://gist.github.com/samalba/6059502
func AssertEqual(t *testing.T, a interface{}, b interface{}, message string) {
  if a == b {
    return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}