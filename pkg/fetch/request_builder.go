package fetch

import (
  "errors"
  "net/http"
  "io/ioutil"
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
    return "", errors.New("http get request failed")
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", errors.New("parsing body failed")
  }

  return string(body), nil
}
