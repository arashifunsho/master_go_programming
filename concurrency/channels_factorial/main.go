package main

import (
	"fmt"
	"strings"
)

func factorial(n int, c chan int) {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}

	c <- f
}

func main() {

	ch := make(chan int)
	defer close(ch)

	go factorial(5, ch)

	f := <-ch

	fmt.Println(f)

	//spawning 20 go routines that calculates factorial using the factorial function
	for i := 1; i <= 20; i++ {
		go factorial(i, ch)
		f := <-ch
		fmt.Println(f)
	}

	fmt.Println(strings.Repeat("#", 10))
	//spawning 10 go routines that calculates factorial using an anonymous function
	for i := 5; i <= 15; i++ {
		go func(n int, c chan int) {
			f := 1
			for i := 2; i <= n; i++ {
				f *= i
			}

			c <- f
		}(i, ch) //calling the anonymous function
		fmt.Printf("Factorial of %d is %d\n", i, <-ch)
	}
}
