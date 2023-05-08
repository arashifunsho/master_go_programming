package main

import (
	"fmt"
	"time"
)

type names []string

// declaring a method (receiver function) on names type
func (n names) print() {
	for i, name := range n {
		fmt.Println(i, name)
	}
}

func main() {
	const day = 24 * time.Hour
	fmt.Printf("%T\n", day)

	seconds := day.Seconds()
	fmt.Printf("%T\n", seconds)
	fmt.Printf("Seconds in a day %v\n", seconds)

	friends := names{"Dan", "Mary"}
	friends.print() //method 1 of calling receiver function

	names.print(friends) //method 2 of calling receiver function

	var n int64 = 9377383844943

	fmt.Println(n)
	fmt.Println(time.Duration(n))
}
