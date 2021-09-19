package main

import (
	"github.com/dreese33/genomics/pkg/fetch"
	"github.com/dreese33/genomics/pkg/fetch/pubchem"
	"github.com/dreese33/genomics/pkg/fetch/pubchem/pubchem_dirs"
)

var names = []string{
	"Aspirin",
}

// The directory for testing the domain substance by name
var domainQueryObj = pubchem.NewPubchemDirectory(
	pubchem_dirs.Substance,
	pubchem_dirs.Name,
)
var domainQuery = domainQueryObj.ConstructDirectory()

// TestEndpoint enumeration
type TestEndpoint int32

const (
	SimpleEndpointTest TestEndpoint = iota
	SubstanceDomainTest
)

// TestEndpoints maps test names to associated endpoints
var TestEndpoints = map[TestEndpoint]fetch.Endpoint{
	SimpleEndpointTest: fetch.NewEndpoint(
		pubchem.BaseURL,
		fetch.JSON,
		pubchem.TestURL,
		pubchem.GetCompoundByName,
		names[0],
	),

	// Domain substance query test
	SubstanceDomainTest: fetch.NewEndpoint(
		pubchem.BaseURL,
		fetch.JSON,
		pubchem.TestURL,
		domainQuery, // directory
		names[0],
	),
}
