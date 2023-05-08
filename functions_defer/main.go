package main

import "fmt"

func foo() {
	fmt.Println("This is foo()")
}

func bar() {
	fmt.Println("This is bar()")
}

func foobar() {
	fmt.Println("This is foobar()")
}

// GO functions can take functions as parameters as well as return functions as values
func increment(x int) func() int {
	return func() int {
		x++
		return x
	}
}

func main() {
	defer foo()
	bar()
	fmt.Println("Just a string after defering foo() and calling bar()")
	defer foobar()

	//anonymous function
	func(msg string) {
		fmt.Println(msg)
	}("I'm an Anonymous function")

	//first class function
	a := increment(10) //a is not of type func() int
	fmt.Printf("%T\n", a)
	fmt.Println(a()) //getting the value of a, we call as function
	fmt.Println(a())

}
