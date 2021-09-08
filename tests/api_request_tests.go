package tests

import (
	"errors"
	"testing"

	"github.com/dreese33/genomics/pkg/fetch"
)

// TestURLExists constructs a url and tests if it exists
func TestURLExists(testEndpoint fetch.Endpoint, testName string, t *testing.T) {
  _, err := fetch.RequestURL(testEndpoint.ConstructURL(), fetch.HEAD)
  if (err != nil) {
    TestFailed(err, t)
    return
  }

  TestPassed(testName)
}


// TestGetRequest tests the HTTP GET request to the provided URL
func TestGetRequest(testEndpoint fetch.Endpoint, testName string, t *testing.T) {
  body, err := fetch.RequestURL(testEndpoint.ConstructURL(), fetch.GET)
  if (len(body) == 0) {
    TestFailed(errors.New("GET request should have a body of at least length 1 "), t)
  }

  if (err != nil) {
    TestFailed(err, t)
    return
  }

  TestPassed(testName)
}