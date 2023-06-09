package main

import (
	"fmt"
)

type km float64
type mile float64

func main() {
	type speed uint

	var s1 speed = 10
	var s2 speed = 20

	fmt.Println(s2 - s1)

	var x uint
	//x = s1 //error, different types
	x = uint(s1) //must convert to use in arthmetic calcs
	_ = x

	//var s3 speed = x //error must convert to type speed before assigning
	var s3 speed = speed(x)
	_ = s3

	var parisToLondon km = 465
	var distanceInMiles mile

	distanceInMiles = mile(parisToLondon) / 0.621
	fmt.Println(distanceInMiles)

}
