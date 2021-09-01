package main

import (
	"fmt"
	"github.com/dreese33/genomics/pkg/fetch"
	"github.com/dreese33/genomics/pkg/fetch/pubchem"
)

func main() {
  const compound = "Aspirin"

  var endpoint = fetch.NewEndpoint(
    pubchem.BaseURL,
    0,
    pubchem.TestURL,
    pubchem.GetCompoundByName,
    compound,
  )
 
  body, err := fetch.RequestURL(endpoint.ConstructURL())
  if err != nil {
    fmt.Println("error occurred")
    fmt.Println(err)
  } else {
    fmt.Println(body)
  }
}
