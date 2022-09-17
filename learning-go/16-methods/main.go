package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Methods or functions with structs in GO")
	gola := User{"Gola", "Gola@gmail.com", true, 25}

	gola.GetStatus()

	// This function doesn't change the actual object
	gola.setEmail()

	// Above function changed email for the copy of "gola" sent to that function not the original one.
	// to verify that
	fmt.Println(gola)
}

// Like we do in any other OOL like java or python, A function defined in a class or to manipulate the
// object is termed as methods. Since we don't have objects or classes in GO, it utilizes structs themselves
// as objects.

// Syntax "func (temp_var Struct_name) function_name(param if any) return_type {}"
func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

// The way we declared this method here, it does not perform operation on the actual object as, in this
// a copy of actually object is being stored in that "u" or temporary variable.
// You can say, as of now we are more or less using "pass by value" type of parameters here.
func (u User) setEmail() {
	u.Email = "test@gmail.com"
	fmt.Println("Email of this user is: ", u.Email)
}
