package main

import (
  "testing"

	"github.com/dreese33/genomics/pkg/fetch"
	"github.com/dreese33/genomics/pkg/fetch/pubchem"
  "github.com/dreese33/genomics/tests"
)

var names = []string{
  "Aspirin",
}

var testEndpoints = []fetch.Endpoint{
  fetch.NewEndpoint(
    pubchem.BaseURL,
    fetch.JSON,
    pubchem.TestURL,
    pubchem.GetCompoundByName,
    names[0],
  ),
}

func TestCompoundByNameURLExists(t *testing.T) {
  tests.TestURLExists(testEndpoints[0], "TestCompoundByNameURLExists", t)
}

func TestGetCompoundByName(t *testing.T) {
  tests.TestGetRequest(testEndpoints[0], "TestGetCompoundByName", t)
}

func TestEndpointTestURL(t *testing.T) {
  body, err := fetch.RequestURL(testEndpoints[0].GetTestURL(), fetch.GET)
  tests.TestBody(t, "TestEndpointTestURL", body, err)
}
