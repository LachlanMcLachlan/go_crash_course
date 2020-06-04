package main

import "fmt"

func main() {
	ids := []int{33, 44, 55, 67, 77}

	// loop through ids using range
	for i, id := range ids {
		fmt.Printf("%d - ID: %d \n", i, id)
	}

	// loop through ids using range - use _ if not using index otherwise will complain
	for _, id := range ids {
		fmt.Printf("ID: %d \n", id)
	}

	// Add ids together
	sum := 0

	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum ", sum)

	// range with map
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s \n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
