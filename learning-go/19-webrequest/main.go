package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Defining URL up here
const url = "https://lco.dev"

func main() {
	fmt.Println("Handling web requests in Go")

	// For http requests we use http package which is a part of netlib module
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Response : ", response.Body)
	}

	// None of the read or body functions closes connection.
	// Use Body.Close() to close the http connection.
	defer response.Body.Close()

	// Data in response.Body is in Bytes format
	data_in_bytes, err := ioutil.ReadAll(response.Body)
	errNilCheck(err)

	fmt.Println("Response : ", string(data_in_bytes))

}

func errNilCheck(err error) {
	if err != nil {
		panic(err)
	}
}
