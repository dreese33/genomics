package pubchem

import (
	"github.com/dreese33/genomics/pkg/fetch"
)

// GetCompoundByName - Constructs url then searches Pubchem api
// for the compound specified
func GetCompoundByName(endpt *fetch.EndpointImpl, name string) (string, error) {
  var url = endpt.GetURL() + "compound/name/" + name + "/" + fetch.Responses[endpt.GetFmt()]
  
  body, err := fetch.RequestURL(url)
  if err != nil {
    return "", err
  }

  return body, nil
}