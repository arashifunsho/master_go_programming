package main

import (
	"fmt"
	"strings"
)

// variadic function declaration //takes 0 or more input params
func f1(a ...int) {
	fmt.Printf("%T\n", a)
	fmt.Printf("%#v\n", a)
}

// modifying the value of the input param via the function
func f2(a ...int) {
	a[0] = 50
}

func sumAndProduct(a ...float64) (float64, float64) {
	var sum float64 = 0
	var product float64 = 1

	for _, v := range a {
		sum += v
		product *= v
	}
	return sum, product
}

// mixing variadic and non-variadic params in func
func personInformation(age int, names ...string) string {
	fullName := strings.Join(names, " ")
	return fmt.Sprintf("Age: %d, Full Name: %s", age, fullName)
}

func main() {
	f1(1, 2, 3, 4)
	f1()
	nums := []int{1, 2}
	nums = append(nums, 3, 4, 5)
	fmt.Println(nums)

	//passing slice to a variadic function
	f1(nums...)
	f2(nums...)
	fmt.Println("Nums: ", nums)

	s, p := sumAndProduct(2.5, 5.9, 10)

	fmt.Println(s, p)

	info := personInformation(30, "Wolfgang", "Amadeus", "Mozart")
	fmt.Println(info)

}
