package main

import (
	"fmt"
	"unorthodox/numbers"
)

func main() {
	var x uint = 40
	fmt.Printf("%d is even: %t\n", x, numbers.Even(x))

	fmt.Println(numbers.OddAndPrime(25))
	fmt.Println(numbers.OddAndPrime(13))
}
