package main

import (
	"fmt"
)

func main() {
	var a = 4
	var b = 5.2

	a = int(b)
	fmt.Println(a, b)

	//var x int
	//x = "5"
	//in Go, all variables are initialized using default values if unassigned
	var value int
	var price float64
	var name string
	var done bool
	fmt.Println(value, price, name, done)

	//this is a single-line comment
	/*
		This is a multi-line comment
	*/

	radius := 5
	pi := 3.1415

	fmt.Println(float64(radius) * 2 * pi)
}
