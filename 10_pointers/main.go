package main

import "fmt"

// sometimes you want to increase performance by passing around the memory address of data rather than the data itself.

func main() {
	a := 5
	// assigns b to a pointer of a (i.e. where it is stored on the system - its memory address)
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b) // *int is a pointer int.

	// read value from memory address using *
	fmt.Println(*b)

	// we can change value a with pointer b
	*b = 10

	fmt.Println(a)
}
