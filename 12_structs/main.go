package main

import (
	"fmt"
	"strconv"
)

// structs are like classes in go. They can have attributes.
// they can have two types of method:
// 1) pointer receiver: change something
// 2) value receiver: don't change anything

// define person struct

type Person struct {
	// define fields
	fistName, lastName, city, gender string
	age                              int
}

// Greeting method (value receiver, since we don't change anything in the struct)
// func (identifier, stuct) name_of_method, return_value
func (p Person) greet() string {
	return "Hello my name is " + p.fistName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver) - we use * in order to update the values inside the struct
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// init Person using struct
	person1 := Person{
		fistName: "Safa", lastName: "al-Seaidy", city: "Boston",
		gender: "F", age: 26}

	// OR

	person2 := Person{"Mike", "Pearson", "London", "m", 25}

	fmt.Println(person2.fistName)

	// can update fields
	person1.age++
	fmt.Println(person1)

	// update age when person has birthday
	fmt.Println(person1.age)
	person1.hasBirthday()
	fmt.Println(person1.age)

	// update person's lastname if woman and get married
	fmt.Println(person1.greet())
	person1.getMarried("McLachlan")
	fmt.Println(person1.greet())
}
