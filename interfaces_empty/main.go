package main

import "fmt"

type emptyInterface interface {
}
type person struct {
	info interface{}
}

func main() {
	var empty emptyInterface
	empty = 5
	fmt.Println(empty)
	empty = "go"
	fmt.Println(empty)
	empty = []int{4, 5, 6}
	fmt.Println(empty)
	fmt.Println(len(empty.([]int))) //type assertion to use the len function

	you := person{}
	you.info = "Your name"
	fmt.Println(you.info)
	you.info = 40
	fmt.Println(you.info)
	you.info = []float64{5.6, 7.5}
	fmt.Println(you.info)
}
