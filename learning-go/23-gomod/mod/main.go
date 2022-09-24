package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Understanding Go Mod")
	greeter()
	// Here mux is used to route things
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")

	// Running server in Go
	// use http.ListenAndServe(port, r)
	// We can use log.Fatal() for logging errors here
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hello in GO MOD")
}

// This method will create a http handler here
func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang</h1>"))
}

/*
	Various commands used

	# For getting any module through internet
	- go get -u github.com/gorilla/mux

	# For ensuring all the imports are satisfied we use
	# It has to walk the package graph to ensure that modules that it walks through are needed
	# for the build
	# It also removes any package which is not required
	# It also brings back packages that are required
	- go mod tidy

	# To verify checksum of modules downloaded through internet use
	- go mod verify

	# To list modules
	- go list
	- go list all
	- go list -m all

	# To list available versions of any modules
	- go list -m -versions github.com/gorilla/mux

	# To check why module are depended on any other module
	- go mod why github.com/gorilla/mux

	# To edit go.mod file
	- go mod edit -go 1.16
	- go mod edit -module github.com/yuvrajyadav13/mod

	# To bring all the modules in current directory
	- go mod vendor

	# To actually use vendor folder instead of web
	# Now it will first look into vendor folder if not found then web
	- go run -mod=vendor main.go
*/
