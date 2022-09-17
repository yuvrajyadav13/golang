package main

import "fmt"

func main() {
	fmt.Println("Structs in Go")

	gola := User{"Gola", "gola@golang.org", true, 20}
	fmt.Println(gola)
	fmt.Printf("Details : %+v", gola)
	fmt.Printf("Name : %v\nEmail : %v", gola.Name, gola.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
