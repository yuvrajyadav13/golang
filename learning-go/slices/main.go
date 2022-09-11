package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in Go")

	// Slices in go are similar to lists in python. They are similar to arrays in go more like they are arrays only in backend.
	// Slices are declare similar to array just without number of elements it should have.
	var fruitList = []string{}
	fmt.Println(fruitList)
	fmt.Printf("Type when array is consider Slices: %T\n", fruitList)

	// Adding elements can be done by usinf append method. Append method takes a slice and values that you want to append.
	// After appending values to slice, append returns the updated slice
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	// append() function does not modify original slice
	var temp = append(fruitList, "Mango", "Banana")
	fmt.Println("Value in original slice:", fruitList)
	fmt.Println("Value of : ", temp)

	// Slicing operation on slice
	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	// Another method of creating slices with "make"
	// In make function we specify the number of elements that slice is to be initialized with.
	// And thats the reason why it is called a slice not a array
	scores := make([]int, 4)

	scores[0] = 65
	scores[1] = 56
	scores[2] = 15
	scores[3] = 95

	fmt.Println("Values in scores: ", scores)

	// We can append operation on slice created using make as we usually do.
	scores = append(scores, 66, 88, 78)
	fmt.Println("Values in scores: ", scores)

	// For sorting a slice
	sort.Ints(scores)
	fmt.Println("Sorted scores: ", scores)

}
