package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to type conversion in GO")
	fmt.Println("Enter any integer: ")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	// This line will throw an error as "input" will containing the \n as well. And ParseFloat will try to convert "4\n" to float.
	addone, err := strconv.ParseFloat(input, 64)

	if err != nil {
		fmt.Println("Error: ", err, "\n")
	} else {
		fmt.Println(addone)
	}

	// Here we will be using strings package for removing any kind of spaces like " ", "\n", "\t", etc.
	addtwo, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("No err even then if in case: ", err)
	} else {
		fmt.Println("Value after error handling and addition", addtwo+1)
	}
}
