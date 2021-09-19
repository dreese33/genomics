package main

import (
	"testing"

	"github.com/dreese33/genomics/tests/helpers"
)

//***************************
// TEST DOMAIN QUERIES
//***************************
func TestSubstanceDomain(t *testing.T) {
	helpers.TestURLExists(TestEndpoints[SubstanceDomainTest], "TestSubstanceDomain", t)
}
