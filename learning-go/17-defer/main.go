package main

import (
	"fmt"
)

func main() {
	// In GO compilation of the program may be as whole
	// However execution is line by line.
	// So here, in case you need that certain statements should not run as they are written
	// you can use "defer" keyword to execute them at last or just before function reached its closing curly brace
	// Statements without "defer" keyword will execute in the same order as they are written in code.
	defer fmt.Println("This is not heading so should be in last or atleast below heading")
	fmt.Println("Defer Keyword in GO")

	// Defer statements are run in reverse order. Or more like in LIFO pattern. Last in First out.
	// In this function first the statements that are without defer will execute and in same order
	// However the last statement with defer will execute before others and then second last one with defer
	// and so on. And in the end first statement with defer will execute
	UseDefer()
}

// Defer statements are stored in stack and that stack is emptied in the end.
func UseDefer() {
	defer fmt.Println("statement-1")
	defer fmt.Println("statement-2")
	defer fmt.Println("statement-3")
	defer fmt.Println("statement-4")
	fmt.Println("Inside UseDefer function.")
	fmt.Println("Trying to understand working of defer keyword. Last to statements will run as it is")
}
