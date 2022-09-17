package main

import "fmt"

func main() {
	fmt.Println("If Else in Go")
	gola := User{"Gola", 10, ""}

	if gola.Age >= 24 {
		gola.Category = "Adult"
	} else if gola.Age > 16 {
		gola.Category = "Teen"
	} else {
		gola.Category = "Minor"
	}

	fmt.Printf("Details of Gola : %+v\n", gola)

	// You can initializing a variable in directly in if
	if age := 10; age > 24 {
		fmt.Printf("Adult\n")
	} else {
		fmt.Printf("Minor\n")
	}

}

type User struct {
	Name     string
	Age      int
	Category string
}
