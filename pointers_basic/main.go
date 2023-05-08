package main

import (
	"fmt"
)

func main() {
	name := "Funsho"
	fmt.Println(&name)

	var x int = 2
	ptr := &x //pointer to address of x variable
	fmt.Printf("ptr is of type %T with a value of %v and address %p\n", ptr, ptr, &ptr)
	fmt.Printf("address of x is %p\n", &x)

	var ptr1 *float64 //declared a pointer to float64
	_ = ptr1

	p := new(int)

	x = 100
	p = &x

	fmt.Printf("p is of type %T with a value of %v\n", p, p)

	//dereferencing operator
	*p = 90 //equivalent to x = 90

	fmt.Println(x, *p)
	fmt.Println("*p == x:", *p == x)

}
