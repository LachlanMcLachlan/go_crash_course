package main

import "fmt"

// maps are key-value pairs
func main() {
	// Define map: `name_of_map := make(map[key_datatype]values_datatype)`
	emails := make(map[string]string)
	// assign kv pairs
	emails["Bob"] = "bob@gmail.com"
	emails["Peter"] = "peter@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])

	// Delete a key
	delete(emails, "Bob")
	fmt.Println(emails)

	// declare map and add kv pairs at the same time
	otherEmails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}
	fmt.Println(otherEmails)
}
