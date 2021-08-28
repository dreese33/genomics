package main

import (
	"fmt"
	"github.com/dreese33/genomics/fetch"
)

func main() {
	respBody, err := fetch.RequestURL(fetch.PubchemAPITest)
	if (err == nil) {
		fmt.Println(respBody)
	}
}
