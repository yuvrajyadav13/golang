package main

import "fmt"

func main() {
	fmt.Println("Loops in Go")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	// In go we don't have multiple keywords to specify different kinds of loop.
	// We have only one keyword "for" for every loop.
	// Kind of loop depends on what type declarations you are using.
	// Below example is of simple for loop with three parameters initializer, condition, increment/decrement
	// In go prefix operators we just have postfix one like var++ or var--
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	// A variant of for-each loop however in this one, i will be holding index not the actual value.
	for i := range days {
		fmt.Println(days[i])
	}

	// Usually variant of for-each loop in this case index will be holding index and day will be holding the value.
	// We can use "continue" and "break" statements as well
	for index, day := range days {
		if index == 3 {
			fmt.Println("Skipping")
			continue
		} else if index == 4 {
			fmt.Println("Breaking")
			break
		} else {
			fmt.Println("Moving forward")
		}
		fmt.Printf("Index is %v and value is %v\n", index, day)
	}

	tempVal := 1

	// This if how we declare while loop in go
	for tempVal < 10 {
		fmt.Println("The value is : ", tempVal)
		if tempVal == 5 {
			// Here goto statement is used to break the flow and go to that label
			goto lco
		}
		tempVal++
	}

	// Goto statements in go
	// Labels are declared with label_name: and then what they should do
lco:
	fmt.Println("Using goto in Go")
}
