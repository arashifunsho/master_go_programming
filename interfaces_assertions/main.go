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

type shape interface {
	area() float64
	perimeter() float64
}

func printShape(s shape) {
	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %v\n", s.area())
	fmt.Printf("Perimeter: %v\n", s.perimeter())
}
func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func main() {
	var s shape = circle{radius: 2.5}
	fmt.Printf("%T\n", s)

	//s.volume() cannot access volume() because it is not defined in the interface
	ball, ok := s.(circle) //using type assertion to get the underlying methods not defined in the interface
	if ok {                //ok holds bool value that tells if type assertion is successful or not
		fmt.Printf("Ball Volume: %v\n", ball.volume())
	}

	//s = rectangle{width: 3.4, height: 2.2}

	//type switch
	// it permits several type assertions in series
	switch value := s.(type) {
	case circle:
		fmt.Printf("%#v has circle type\n", value)
	case rectangle:
		fmt.Printf("%#v has rectangle type\n", value)
	}

}
