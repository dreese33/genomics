package main

import (
	"fmt"
	"testing"

	"github.com/dreese33/genomics/pkg/fetch"
	"github.com/dreese33/genomics/pkg/fetch/pubchem"
)

var names = []string{
  "Aspirin",
}

var testEndpoints = []fetch.Endpoint{
  fetch.NewEndpoint(
    pubchem.BaseURL,
    fetch.JSON,
    pubchem.TestURL,
    pubchem.GetCompoundByCID,
    names[0],
  ),
}

// TestURLExists constructs a url and tests if it exists
func TestURLExists(t *testing.T) {
  _, err := fetch.RequestURL(testEndpoints[0].ConstructURL(), fetch.HEAD)
  if (err != nil) {
    t.Errorf("Error occurred fetching URL %s", err.Error())
  }

  fmt.Println("TestURLExists -- Test Passed")
}


// TestURLRequest tests the HTTP GET request to the provided URL
func TestURLRequest(t *testing.T) {
  _, err := fetch.RequestURL(testEndpoints[0].ConstructURL(), fetch.GET)
  if (err != nil) {
    t.Errorf("Error occurred getting data %s", err.Error());
  }

  fmt.Println("TestURLERequest -- Test Passed")
}