package main

import "fmt"

func main() {
	fmt.Println("Maps in Go")

	// Maps in go can be declared like this
	groceries := make(map[string]string)

	// For adding any key value pairs in a map
	groceries["Peanuts"] = "500gm"
	groceries["Flour"] = "5kg"
	groceries["lentils"] = "1kg"

	fmt.Println("Map : ", groceries)
	fmt.Println("Peanuts: ", groceries["Peanuts"])

	// for deleting any value from map we use delete function
	delete(groceries, "Flour")
	fmt.Println("Map: ", groceries)

	for key, value := range groceries {
		fmt.Printf("Key = %s, Value = %s\n", key, value)
	}
}
