package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	hello := "Hello"
	fmt.Println(hello)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter \"GO\" :")

	// Comm OK error syntax
	// In Comma OK Error syntax, var before comma is treated as try and and var after comma is treated as catch
	// Not exactly same but. in below case whatever input user gives will be storage in "input" and if any error is encountered that
	// will be stored in _. we can replace "_" with any var name. As "_" is used when it doesn't matter whatever the value that will be storage in error part.
	// "_" can be used in place or input as well if whatever the input user provide doesn't matter.

	// In reader.ReadString('\n') "\n" means that we are telling reader to read a string until it encounter a newline.
	// You can use any character you want. It will add that character into your input as well.
	input, _ := reader.ReadString('\n')
	fmt.Println("Welcome in world of ", input)
	fmt.Printf("Type of input : %T", input)
}
