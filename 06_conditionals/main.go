package main

import "fmt"

func main() {
	x := 5
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d \n", x, y)
	} else {
		fmt.Printf("%d is less than %d \n", y, x)
	}

	// else if
	color := "blue"

	if color == "red" {
		fmt.Println("colour is red")
	} else if color == "blue" {
		fmt.Println("colour is blue")
	} else {
		fmt.Println("not sure what colour this is...")
	}

	// Switch
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is neither red nor blue.")
	}

}
