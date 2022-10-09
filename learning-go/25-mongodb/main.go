package main

import (
	"25-mongodb/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Working with mongo db in Go")
	fmt.Println("Server is getting started")

	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening")

}

// mongodb+srv://ydvyuvi:<password>@cluster0.pzpbhhb.mongodb.net/?retryWrites=true&w=majority
