package main

import (
	"fmt"
	"math"
)

// function with eno params
func f1() {
	fmt.Println("This is f1() function.")
}

// function with input params
func f2(a int, b int) {
	fmt.Println("Sum: ", a+b)
}

// funtion with multiple, similar typed input params
func f3(a, b, c int, d, e float64, s string) { //shorthand function declaration. variables of same type can have just one type declaration
	fmt.Println(a, b, c, d, e, s)
}

// function that returns single value
func f4(a float64) float64 {
	return math.Pow(a, a)
}

// function that returns multiple values
func f5(a, b int) (int, int) {
	return a + b, a * b
}

// function with named return value
func sum(a, b int) (s int) {
	fmt.Println(s) //s is initially 0 at this point

	s = a + b
	return //naked return - we  can return a named parameter without adding it to the return statement
}

func main() {
	f1()
	f2(5, 7)
	f3(4, 5, 6, 4.4, 7.8, "golang")
	p := f4(5.1)
	fmt.Println(p)

	a, b := f5(8, 8)
	fmt.Println(a, b)

	mySum := sum(4, 8)
	fmt.Println(mySum)
}
