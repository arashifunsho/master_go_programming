package main

import (
	"fmt"
	"math"
)

func main() {
	var x uint8 = 255
	x++            //overflow
	fmt.Println(x) //x resets to 0

	//a := int8(255 + 1)
	//fmt.Println(a)
	var b int8 = 127
	fmt.Printf("%d\n", b+1) //b overflows and returns to the type's minimum value
	b = -128
	b--
	fmt.Printf("%d\n", b)

	//float overflow

	f := float32(math.MaxFloat32)
	fmt.Println(f)

	f *= 1.2 //f overflows to infinite
	fmt.Println(f)
}
