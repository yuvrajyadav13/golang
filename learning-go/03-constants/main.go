package main

import "fmt"

// Declaring constant values
const username = "hello"

// Go treats variable with first literal in upper case a public variable
const Publicvar = "world"

func main() {
	fmt.Println(username)
	var username = "world"
	fmt.Println(username)
	fmt.Println(Publicvar)
}
