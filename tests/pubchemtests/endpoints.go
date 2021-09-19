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

var testEndpoints = []fetch.Endpoint{
	fetch.NewEndpoint(
		pubchem.BaseURL,
		fetch.JSON,
		pubchem.TestURL,
		pubchem.GetCompoundByName,
		names[0],
	),

	// Domain substance query test by name
	fetch.NewEndpoint(
		pubchem.BaseURL,
		fetch.JSON,
		pubchem.TestURL,
		domainQuery, // directory
		names[0],
	),
}
