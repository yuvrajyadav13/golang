package main

import "fmt"

func main() {
	fmt.Println("Pointers in Golang")

	// Pointers are same as in any other language
	var ptr *int

	// By Default a pointer when not initialized holds "nil"
	fmt.Println("Default value : ", ptr)
	fmt.Printf("Type: %T \n", ptr)

	myNum := 23

	// As in C, & can be used to get memory address of any variable
	var ptr2 = &myNum
	fmt.Println("Address of myNum : ", &myNum)
	fmt.Println("Value in myNum : ", myNum)
	fmt.Println("Value in ptr : ", ptr2)
	fmt.Println("Value in ptr is referencing : ", *ptr2)

	// Performing any actions using pointers will directly affect the value it is referenced to.
	*ptr2 = *ptr2 * 2
	fmt.Println("Value in myNum : ", myNum)

	// **ptr, double ** like in C are not supported in Go
	fmt.Println(ptr2)
}
