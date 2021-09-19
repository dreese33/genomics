package main

import (
	"testing"

	"github.com/dreese33/genomics/pkg/fetch"
	"github.com/dreese33/genomics/tests/helpers"
)

func TestCompoundByNameURLExists(t *testing.T) {
	helpers.TestURLExists(TestEndpoints[SimpleEndpointTest], "TestCompoundByNameURLExists", t)
}

func TestGetCompoundByName(t *testing.T) {
	helpers.TestGetRequest(TestEndpoints[SimpleEndpointTest], "TestGetCompoundByName", t)
}

func TestEndpointTestURL(t *testing.T) {
	body, err := fetch.RequestURL(TestEndpoints[SimpleEndpointTest].GetTestURL(), fetch.GET)
	helpers.TestBody(t, "TestEndpointTestURL", body, err)
}
