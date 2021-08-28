package fetch

import (
	"errors"
	"net/http"
	"io/ioutil"
)

/*
RequestURL returns the url response body
in the specified format
*/ 
func RequestURL(url string) (string, error) {
	resp, err := http.Get(url)
	if (err != nil) {
		return "", errors.New("http get request failed")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if (err != nil) {
		return "", errors.New("parsing body failed")
	}

	return string(body), nil
}
