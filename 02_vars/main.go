package main

import "fmt"

// can make global variables
var middleName string = "Kennedy"

func main() {
	// this is a comment
	// structure is 'type_of_thing' 'name_of_thing' '(optional)data_type' = 'thing'
	var name string = "Lachlan"
	// go prefers camelCase
	var otherName = "Stuart"
	var age int = 26
	var isCool = true

	// if you don't want to be able to change something later in the program,
	// use const instead of var
	const isAwsome = false

	// shorthand variable assignment (works only inside functions)
	brother := "Jacob"
	multiple, variables := "hello", "world"

	// you have to use variables in go - you cannot declare them and then not use them
	fmt.Println(name, otherName, middleName, age, isCool, isAwsome, brother, "blah", multiple, variables)
	// print the type of a variable
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
}
