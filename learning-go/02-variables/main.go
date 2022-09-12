package main

import "fmt"

func main() {
	var name string = "Hello"
	fmt.Println(name)
	fmt.Printf("Type: %T \n\n", name)

	// int can be used as int, int 8, int16, int32, int64. if used int, it can be either int32 or int64.
	var age int = 20
	fmt.Println(age)
	fmt.Printf("Type: %T \n\n", age)

	var islogged bool = true
	fmt.Println(islogged)
	fmt.Printf("Type: %T \n\n", islogged)

	// float can be used as float32, float64. float32 provides precision upto 5 decimal places while float64 can provide upto 13 decimal places
	var weight float32 = 50.02
	fmt.Println(weight)
	fmt.Printf("Type: %T \n\n", weight)

	// uint or unsigned int can be used as uint, uint8, unit16, unit32, unit64.
	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("Type: %T \n\n", smallval)

	var midval uint32 = 256
	fmt.Println(midval)
	fmt.Printf("Type: %T \n\n", midval)

	// If no value is provided in variables they are empty int initialized to zero and string is empty string, bool will be
	var non_init_int int
	var non_init_bool bool
	var non_init_string string
	fmt.Printf("%T: %d, %T: %b, %T: %s \n\n", non_init_int, non_init_int, non_init_bool, non_init_bool, non_init_string, non_init_string)

	// If variable type is not provided go we define it according to its value
	var wo_type_int = 1
	var wo_type_bool = true
	var wo_type_string = "Hello"
	fmt.Printf("%T, %T, %T \n\n", wo_type_int, wo_type_bool, wo_type_string)

	// Without var style. By using WALRUS OPERATOR. Walrus operators are not allowed outside methods/func like for global declarations.
	walrus_var := "world"
	fmt.Printf("%T: %s \n", walrus_var, walrus_var)

}
