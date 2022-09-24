package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=dfghh435bgf"

func main() {
	fmt.Println("URL handling in Go")
	fmt.Println("Url: ", myurl)

	// url library in Go lang helps in working with URLs
	// Below method parses url and degenerates it into varies parts that a url contains
	result, _ := url.Parse(myurl)

	// Schema defines like which protocol that url is using in above case it is https
	fmt.Println(result.Scheme)

	// Host is that website base url like google.com. It includes hostname and port.
	// hostname:port
	fmt.Println(result.Host)

	// Path specifies the path to the webpage that you are trying to acccess
	fmt.Println(result.Path)

	// Port specifies through which port you are accessing that url or like which port is getting used while accessing that url
	fmt.Println(result.Port())

	// Raw query gives you raw version of query that you have used in the url
	// example : schema://host:port/path?rawquery
	fmt.Println(result.RawQuery)

	// Hostname is actual base url without port
	fmt.Println(result.Hostname())

	// Query return rawQuery or parameters in formated way or in a map
	fmt.Println(result.Query())

	for _, val := range result.Query() {
		fmt.Println(val)
	}

	// Defining a URL
	// While defining url,always pass address of that object not its copy.
	partsofurl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=gola",
	}

	fmt.Println(partsofurl.String())
}
