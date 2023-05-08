package main

import (
	"fmt"
	"strings"
)

func main() {
	var numbers [4]int //array initialized to zeros by default

	fmt.Printf("%d\n", numbers)
	fmt.Printf("%#value\n", numbers) //another way to print arrays

	var a1 = [4]float64{} //array declaration
	fmt.Printf("%#value\n", a1)

	var a2 = [3]int{-10, 1, 100} //array declaration and initialization
	fmt.Printf("%#value\n", a2)

	a3 := [4]string{"Dan", "Diana", "Paul", "John"} //short declaration for arrays
	fmt.Printf("%#value\n", a3)

	a4 := [4]string{"x", "y"} //arrays can be initialized with partial values
	fmt.Printf("%#value\n", a4)

	//ellipsis operator
	a5 := [...]int{1, 2, 3, 5, -10, 15} //ellipsis operator automatically counts the number of elements as declared in the array
	fmt.Printf("The length of a5 is %d\n", len(a5))

	a6 := [...]int{1, //multi-line array declaration
		2,
		3,
		4,
		5, //this comma is mandatory
	}
	fmt.Println(a6)

	numbers1 := [3]int{10, 20, 30}
	fmt.Printf("%#value\n", numbers1)

	numbers1[0] = 7
	fmt.Printf("%#value\n", numbers1)
	//numbers1[5] = 100 //out of bounds error

	//iterating over an array
	for index, value := range numbers1 { //method 1
		fmt.Println("index:", index, "value:", value)
	}
	fmt.Println(strings.Repeat("#", 20))
	for i := 0; i < len(numbers1); i++ { //method 2
		fmt.Println("index:", i, "value:", numbers1[i])
	}
	fmt.Println(strings.Repeat("#", 20))

	//multi-dimensional array
	balances := [2][3]int{
		[3]int{5, 6, 7},
		{8, 9, 10}, //the type declaration is optionals
	}
	fmt.Println(balances)

	m := [3]int{1, 2, 3}
	n := m //creates a copy of the array

	fmt.Println("n is equal to m:", n == m) //arrays are equal
	m[0] = -1
	fmt.Println("n is equal to m:", n == m) //arrays no longer equal as one was modified

	//arrays with keyed elements
	grades := [3]int{
		1: 10,
		0: 5, //keyed elements can be in any order
		2: 7,
	}
	fmt.Println(grades)

	accounts := [3]int{2: 50}
	fmt.Println(accounts)

	names := [...]string{
		5: "Dan",
	}
	fmt.Println(names, len(names))

	cities := [...]string{
		5:        "Paris",
		"London", //unkeyed element gets its index from the last keyed element
		1:        "NYC",
	}
	fmt.Printf("%#v\n", cities)

	weekend := [7]bool{5: true, 6: true} //weekdays are initialized with the default, false
	fmt.Println(weekend)

}
