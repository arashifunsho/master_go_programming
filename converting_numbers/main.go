package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x = 3
	var y = 3.1
	x *= int(y)
	fmt.Println(x)

	fmt.Printf("%T\n", y)

	x = int(float64(x) * y)
	fmt.Println(x)

	s := string(99)
	fmt.Println(s)

	var myStr = fmt.Sprintf("%f", 44.2)
	fmt.Println(myStr)

	s1 := "3.123"
	fmt.Printf("%T\n", s1)

	var f1, _ = strconv.ParseFloat(s1, 64)
	fmt.Println(f1 * 1.0)

	i, _ := strconv.Atoi("-50")
	s2 := strconv.Itoa(20)
	fmt.Printf("i type is %T, i value is %v\n", i, i)
	fmt.Printf("s2 type is %T, i value is %q\n", s2, s2)
}
