package main

import "fmt"

func main() {

	// long method
	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++ // (same as `i = i + 1`)
	}

	// short method: assignment, condition, increment {
	//	body
	//}

	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d \n", i)
	}

	// fizz buzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
