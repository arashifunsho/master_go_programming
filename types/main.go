package main

import "fmt"

func main() {
	var r rune = 'f'
	fmt.Printf("%T\n", r)

	//string typr

	//var s string = "Hello Go!"
	//GO ARRAY DECLARATION
	var numbers = [4]int{4, 5, -5, 100} //static size
	fmt.Printf("%T\n", numbers)

	//GO SLICE DECLARATION
	var cities = []string{"London", "Tokyo", "New York"} //size can grow/shrink
	fmt.Printf("%T\n", cities)

	//GO MAP DECLARATION
	balances := map[string]float64{
		"USD": 2332.2,
		"EUR": 511.1,
	}
	fmt.Printf("%T\n", balances)

	//GO STRUCT DECLARATION
	type Person struct {
		name string
		age  int
	}
	var you Person
	fmt.Printf("%T\n", you)

	//GO POINTER DECLARATION
	var x int = 1
	ptr := &x
	fmt.Printf("%T\t%v\n", ptr, ptr)

	fmt.Printf("%T\n", f)
}

// GO FUNCTION TYPE
func f() {

}
