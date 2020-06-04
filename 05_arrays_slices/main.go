package main

import "fmt"

func main() {
	// Arrays - must be fixed type and length (declared here)
	var fruitArr [2]string

	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// declare and assign array at the same time
	anotherArr := [2]string{"Grapefruit", "Banana"}

	// slices (don't have to assign length)
	fruitSlice := []string{"Apple", "Passion_Fruit", "Kiwi"}

	fmt.Println(fruitArr)
	fmt.Println(anotherArr)
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[0:2])
}
