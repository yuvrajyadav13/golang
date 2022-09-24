package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("File handling in Go")
	content := "Test text for storing in file"

	// Here we are using os package for creating a file. Like in any other language.
	// os package has a lot of things like permission related commands and other things so
	// do look into this.
	// os.create creates a new file if it is not there. And replaces it if it is already there.
	file, err := os.Create("./test.txt")

	if err != nil {
		// Here panic is similar to throw statement in java incase of exceptions.
		// panic will throw the error on console and write the error it received
		panic(err)
	} else {

		// Here io.WriteString function writes string to the file. first argument is file and second is
		// the string or content that is to be added in that file
		length, err := io.WriteString(file, content)

		if err != nil {
			panic(err)
		} else {
			fmt.Println(length)
		}
	}

	defer file.Close()
	fmt.Println(file.Name())
	readFile(file.Name())
}

func readFile(filename string) {
	// ioutil is now deprecated. Similar fuctionality is provided by io and os packages now.
	// ioutil.ReadFile needs path to the file. Relative path can also be given.
	// It returns data in bytes format. Hence it has to be converted to string first.
	// string() can be used for that. There are other functions as well.
	byte_format_data, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println("Byte format data that ioutil sent : ", byte_format_data)
	fmt.Println("Human readable way: ", string(byte_format_data))
}
