package tests

import (
  "testing"
  "fmt"
)

// AssertEqual ensures that two types are equal
func AssertEqual(t *testing.T, a interface{}, b interface{}, message string) {
  if a == b {
    return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}