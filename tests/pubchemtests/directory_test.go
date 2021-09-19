package main

import (
	"testing"

	"github.com/dreese33/genomics/tests/helpers"
)

//***************************
// TEST DOMAIN QUERIES
//***************************
func TestDomainByName(t *testing.T) {
	helpers.TestGetRequest(testEndpoints[1], "TestDomainByName", t)
}
