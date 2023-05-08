package main

import "fmt"

func main() {
	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)

	fmt.Println(c1, c2, c3)

	const (
		c11 = iota
		c22
		c33
	)

	fmt.Println(c11, c22, c33)

	const (
		North = iota
		East
		South
		West
	)

	fmt.Println(West)

	const (
		a = iota * 2
		b
		c
	)

	fmt.Println(a, b, c)

	//x=-2, y=-4, z=-5
	const (
		x = -(iota + 2)
		_ //skipping this value initializer using the blank identifier
		y
		z
	)
	fmt.Println(x, y, z)
}
