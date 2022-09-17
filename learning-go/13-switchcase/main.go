package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch-Case in Go")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	// Switch Case in go are similar to other languages. Only thing different in go's switch case
	// is we don't have to explicitly mention break statement. It will run only those case which it
	// encounters. There is no automatic fallthrough.
	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spots")
	case 4:
		fmt.Println("You can move 4 spots")
	case 5:
		fmt.Println("You can move 5 spots")
	case 6:
		fmt.Println("You can move 6 spots and roll again")
	default:
		fmt.Println("Retry")
	}

	// In case you require fallthrough. You have to mention fallthrough explicitly in case you need.
	// fallthrough should be explicitly specified in each case if you require that every case after
	// what you encountered, must ran.

	// If you specify fallthrough in last case then default case will also run.
	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
		fallthrough
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spots")
		fallthrough
	case 4:
		fmt.Println("You can move 4 spots")
		fallthrough
	case 5:
		fmt.Println("You can move 5 spots")
		fallthrough
	case 6:
		fmt.Println("You can move 6 spots and roll again")
	}
}
