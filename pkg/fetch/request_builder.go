package fetch

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// HTTP methods enumeration
type HTTP int32

// Methods for data retrieval/mutation
const (
  GET HTTP = iota
  HEAD
)


// RequestURL returns the url response body
// in the specified format
func RequestURL(url string, method HTTP) (string, error) {
  var resp *http.Response
  var err error
  switch method {
    case GET:
      resp, err = http.Get(url)
    case HEAD:
      resp, err = http.Head(url)
  }

  if err != nil {
    return "", err
  }

  if (resp.StatusCode != 200) {
    return "", errors.New(
      "Failed with status code " + fmt.Sprint(resp.StatusCode),
    )
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", err
  }

  return string(body),nil
}
