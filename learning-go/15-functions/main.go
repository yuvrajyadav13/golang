package main

import "fmt"

func main() {
	fmt.Println("Functions in Go")

	// To call a function. Similar to c/c++
	temp()

	// You cannot define function inside a function
	/*
		func temp2() {
			fmt.Println("Function inside function not allowed")
		}

	*/

	fmt.Println("Addition is : ", addition(6, 6))
	fmt.Println("Sum is : ", alladder(5, 5, 50, 65))

	value, statement := multivaluereturn(5, 6, 7, 5)
	fmt.Println(statement, value)
}

// In go we declare functions using "func" keyword
// Similar to c or c++, execution of program starts from main()

func temp() {
	fmt.Println("Temporary function-1")
}

// Syntax for function is "func function_name(param param_type, param2, param_type ..) return_type {}"
// Function parameters and return_type are known as function signatures
func addition(valOne int, valTwo int) int {
	return valOne + valTwo
}

// function with variable number of arguments or varargs as in python
//
//	It is also called variadic functions
func alladder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}

// Returning mutliple values from functions

func multivaluereturn(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "Addition of all values through multivaluereturn function is : "
}
