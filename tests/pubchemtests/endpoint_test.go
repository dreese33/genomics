package main

import (
	"testing"

	"github.com/dreese33/genomics/pkg/fetch"
	"github.com/dreese33/genomics/tests/helpers"
)

func TestCompoundByNameURLExists(t *testing.T) {
	helpers.TestURLExists(testEndpoints[0], "TestCompoundByNameURLExists", t)
}

func TestGetCompoundByName(t *testing.T) {
	helpers.TestGetRequest(testEndpoints[0], "TestGetCompoundByName", t)
}

func TestEndpointTestURL(t *testing.T) {
	body, err := fetch.RequestURL(testEndpoints[0].GetTestURL(), fetch.GET)
	helpers.TestBody(t, "TestEndpointTestURL", body, err)
}
