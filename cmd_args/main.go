package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//fmt.Println("os.Args", os.Args)
	//fmt.Println("Path: ", os.Args[0])
	//fmt.Println("1st Argument: ", os.Args[1])
	//fmt.Println("2nd Argument: ", os.Args[2])
	//fmt.Println("No. of Items in os.Args: ", len(os.Args))

	var result, err = strconv.ParseFloat(os.Args[1], 64)
	fmt.Printf("Result: %v\nType: %T\n", result, os.Args[1])
	_ = err
}
