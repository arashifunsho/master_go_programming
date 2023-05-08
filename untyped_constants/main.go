package main

import "fmt"

func main() {
	const a float64 = 5.1 //typed constant

	const b = 6.7 //untyped constant

	const c float64 = a * b
	const str = "Hello" + "Go!"

	const d = 5 > 10

	fmt.Println(d)
	//const x int = 5
	//const y float64 = 2.2 * x

	const x = 5
	const y = 2.2 * x
	fmt.Printf("%T\n%T\n", y, x)

	var i int = x     //x changes type to int
	var j float64 = x //same as var j float64 = float64(x)
	var p byte = x

	fmt.Printf("%T\t%T\t%T", i, j, p)
}
