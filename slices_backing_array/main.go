package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s1 := []int{10, 20, 30, 40, 50}
	s3, s4 := s1[0:2], s1[1:3]

	s3[1] = 600 //modifying one slice changes the element in the backing array and affects all slices derived from it
	fmt.Println(s1)
	fmt.Println(s4)

	arr1 := [4]int{10, 20, 30, 40}
	slice1, slice2 := arr1[0:2], arr1[1:3] //arr1 is now the backing array of slice1 and slice2
	arr1[1] = 2
	fmt.Println(arr1, slice1, slice2) //since arr1 was modified, this also modifies all its slices

	//creating a totally new slice using append
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}
	newCars = append(newCars, cars[0:2]...) //unpack the cars slice using ...(ellipsis) and append to newCars
	cars[0] = "Nissan"
	fmt.Println(cars, newCars) //cars and newCars do not share the same backing array

	s := []int{10, 20, 30, 40, 50}
	newSlice := s[0:3]
	fmt.Println(len(newSlice), cap(newSlice)) //cap returns number of element in the slice's backing array

	fmt.Printf("%p\n", &s) //printing the address of s

	fmt.Printf("%p\t%p\n", &s, &newSlice)

	//check memory size
	a := [5]int{1, 2, 3, 4, 5}
	c := []int{1, 2, 3, 4, 5}

	fmt.Printf("array size in byte: %d\n", unsafe.Sizeof(a))
	fmt.Printf("slices size in byte: %d\n", unsafe.Sizeof(c))
}
