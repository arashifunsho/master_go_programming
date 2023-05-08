package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

//func printCircle(c circle) {
//	fmt.Println("Shape:", c)
//	fmt.Println("Area:", c.area())
//	fmt.Println("Perimeter:", c.perimeter())
//}
//func printRectangle(r rectangle) {
//	fmt.Println("Shape:", r)
//	fmt.Println("Area:", r.area())
//	fmt.Println("Perimeter:", r.perimeter())
//
//}

// refactoring and using interfaces
// an interface contains only the signatures of the methods, but not their implementation
type shape interface {
	area() float64
	perimeter() float64
}

// any type that implements the interface is also of type of the interface
// rectangle and circle values are also of type shape
func printShape(s shape) {
	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %v\n", s.area())
	fmt.Printf("Perimeter: %v\n", s.perimeter())
}

func main() {
	c1 := circle{radius: 5}
	r1 := rectangle{width: 3, height: 2.1}

	//printCircle(c1)
	//printRectangle(r1)
	printShape(c1)
	printShape(r1)

}
