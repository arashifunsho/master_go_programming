package main

import "fmt"

func main() {
	var age = 30
	fmt.Println("Age:", age)
	var name = "Dan"
	fmt.Println("Your name is: ", name)
	//_ = name

	s := "Learning golang!"
	fmt.Println(s)

	car, cost := "Audi", 50000 //declaring multiple variables
	fmt.Println(car, cost)

	car, year := "BMW", 2018
	_ = year

	var opened = false
	opened, file := true, "a.txt" //existing variables can only be redeclared this way if at least one is new
	_, _ = opened, file

	//another method for multiple variable declarations
	var (
		salary    float64
		firstName string
		gender    bool
	)
	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	//multiple assignments
	var i, j int
	i, j = 5, 8 //tuple assignment
	//swapping variables
	i, j = j, i
	fmt.Println(i, j)

	sum := 5 + 2.3
	fmt.Println(sum)
}
