package tests

import (
	"fmt"
	"testing"
)

// TestPassed logs the test case if it passes
func TestPassed(testName string) {
	fmt.Println(testName + " -- Test Passed")
}

// TestFailed logs the test case if it fails
func TestFailed(err error, t *testing.T) {
	t.Errorf("Error occurred %s", err.Error())
}
