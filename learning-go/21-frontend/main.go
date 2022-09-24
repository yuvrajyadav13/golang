package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Perform Get/Post requests in Go")
	// Perform get request
	PerformGetRequest()

	// Perform post request
	PerformPostRequest()

	// Perform postform request
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	// To get what status code was received in request
	fmt.Println("Status code: ", response.StatusCode)
	// Length of content
	fmt.Println("Cpntent length: ", response.ContentLength)

	// Type of response.Body
	fmt.Printf("%T", response.Body)
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

	// Another way of getting content into string
	// Here we are using strings Builder
	// In some cases this is better
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte count : ", byteCount)
	fmt.Println("String is: ", responseString.String())
}

func PerformPostRequest() {
	const myurl = "http://localhost:8000/post"

	// Fake json payload
	// Some json has to be sent in Post request
	requestbody := strings.NewReader(`
		{
			"country":"India",
			"Name":"gola",
			"config":"none"
		}
	`)

	// In Post method, url, content type, and content is necessary
	response, err := http.Post(myurl, "application/json", requestbody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)

	fmt.Println("Content: ", string(content))

}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	// Create formdata

	data := url.Values{}
	data.Add("firstname", "Gola")
	data.Add("lastname", "google")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println("Content: ", string(content))
}
