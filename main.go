package main

import (
	"fmt"
	"github.com/dreese33/genomics/api"
	"github.com/dreese33/genomics/pkg/fetch"
	"github.com/dreese33/genomics/pkg/fetch/pubchem"
)

func main() {
  const compound = "Aspirin"

  var endpoint = fetch.EndpointImpl{
    URL: api.PubchemAPI,
    Fmt: 0,
    TestURL: api.PubchemAPITest,
  }

  body, err := pubchem.GetCompoundByName(&endpoint, compound)
  if err != nil {
    fmt.Println("error occurred")
    fmt.Println(err)
  } else {
    fmt.Println(body)
  }
}
