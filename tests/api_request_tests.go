package tests

import (
	"errors"
	"testing"

	"github.com/dreese33/genomics/pkg/fetch"
)

const EmptyBody = "GET request should have a body of at least length 1 "

// TestURLExists constructs a url and tests if it exists
func TestURLExists(testEndpoint fetch.Endpoint, testName string, t *testing.T) {
  _, err := fetch.RequestURL(testEndpoint.ConstructURL(), fetch.HEAD)
  if (err != nil) {
    TestFailed(err, t)
    return
  }

  TestPassed(testName)
}

// TestBody ensures that the body of the response is of a certain length
func TestBody(t *testing.T, testName string, body string, err error) {
  if (len(body) == 0) {
    TestFailed(errors.New(EmptyBody), t)
    return
  }

  if (err != nil) {
    TestFailed(err, t)
    return
  }

  TestPassed(testName)
}

// TestGetRequest tests the HTTP GET request to the provided URL
func TestGetRequest(testEndpoint fetch.Endpoint, testName string, t *testing.T) {
  body, err := fetch.RequestURL(testEndpoint.ConstructURL(), fetch.GET)
  TestBody(t, testName, body, err)
}
