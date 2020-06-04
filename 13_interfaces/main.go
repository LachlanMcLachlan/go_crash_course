package main

import (
	"fmt"
	"math"
)

// Define interface

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

// value receiver functions to implement area method for the
// Circle and Rectangle structs
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r Rectangle) area() float64 {
	return r.width * r.height
}

// create a method called getArea which takes an interface and calls
// the area() method. In Python you wouldn't have to define the interface
// because the argument received by getArea could be either a Circle or
// a Rectange since you don't have to define a type for the argument.
// Hence static typing forces you to make an interface in order to supply
// a type to this method.
func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectange := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectange))
}
