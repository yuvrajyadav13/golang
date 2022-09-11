package main

import "fmt"

func main() {
	fmt.Println("Arrays in Go")

	// This is one way of declaring Arrays. Arrays are by default initialized with what every default value there base type has.
	// Like a array of ints will be initialize with 0
	// A array of strings will be initialized with " "
	var fruitList [10]string
	var intList [4]int

	fmt.Println(fruitList)
	fmt.Println(intList)

	// Type of Arrays will output [n]datatype. Eg for fruitList => [10]string
	fmt.Printf("%T ", fruitList)

	// As arrays are initialize with there default values at each index. Even if no value is given there default value is always there.
	// Look for 3rd element in output of below list. There is an extra space
	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Pomogranate"

	fmt.Println(fruitList)

	// In floating point arrays 1.0 is outputed as 1 without .0
	var floatList = [3]float32{1.05, 2.07, 3.0}
	fmt.Println(floatList)
}
