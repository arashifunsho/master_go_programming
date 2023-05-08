package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}
type object interface {
	volume() float64
}

// declaring a geometry interface by embedding shape and object interfaces and adding its internal method signature
type geometry interface {
	shape
	object
	getColor() string
}
type cube struct {
	edge  float64
	color string
}

// cube must declare the getColor() method before it is considered as implementing the geometry interface
func (c cube) getColor() string {
	return c.color
}

func (c cube) area() float64 {
	return 6 * math.Pow(c.edge, 2)
}
func (c cube) volume() float64 {
	return math.Pow(c.edge, 3)
}

func measure(g geometry) (float64, float64) {
	return g.area(), g.volume()
}

func main() {
	c := cube{edge: 2}
	a, v := measure(c)

	fmt.Printf("Area:%v, Volume: %v\n", a, v)

}
