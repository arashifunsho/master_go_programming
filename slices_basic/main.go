package main

import "fmt"

func main() {
	var cities []string //initialized empty string slice
	fmt.Println("cities is equal to nil: ", cities == nil)
	fmt.Printf("cities %#v\n", cities)

	fmt.Println(len(cities))
	//cities[0] = "London" //error

	numbers := []int{2, 3, 4, 5} //creating slice using int literals
	fmt.Println(numbers)

	nums := make([]int, 2) //initialized using make function; same as []int{0,0}
	fmt.Printf("%#v\n", nums)
	fmt.Println(nums)

	type names []string
	friend := names{"Dan", "Maria"}
	fmt.Println(friend)

	myFriend := friend[0]

	fmt.Println("My Best friend is", myFriend)
	friend[0] = "Gabriel"
	fmt.Println("My Best friend is", friend[0])

	//iterating over a slice
	for i, v := range numbers {
		fmt.Printf("index: %v\tvalue:%v\n", i, v)
	}

	var n []int
	n = numbers //slices can be assigned to each other if they are of the same type
	fmt.Println(n)

	//comparing slices
	var m []int
	fmt.Println(m == nil) //true

	l := []int{}
	fmt.Println(l == nil) //false

	a, b := []int{1, 2, 3}, []int{1, 2, 3}
	//fmt.Println(a == b) //error cannot conpare two slices using ==

	//comparing slices
	var eq = true
	for i, valueA := range a {
		if valueA != b[i] {
			eq = false
			break
		}
	}
	a = nil
	if len(a) != len(b) { //when comparing slices, if one is nil, the comparison will always return true. Best to also check the sizes
		eq = false
	}
	if eq {
		fmt.Println("a & b slices are equal")
	} else {
		fmt.Println("a & b slices are not equal")
	}

	//appending to and copying  a slice
	numbers1 := []int{2, 3}
	numbers1 = append(numbers1, 10) //appends a new number to the slice and returns a new slice. Here was are overwriting the existing slice variable
	fmt.Println(numbers1)

	numbers1 = append(numbers1, 20, 30, 40) //append is also a variadic function. It can take multiple elements
	fmt.Println(numbers1)

	//appending one slice to another
	nn := []int{100, 200}
	numbers1 = append(numbers1, nn...) //... lists all elements in nn
	fmt.Println(numbers1)

	src := []int{10, 20, 30}
	dest := make([]int, len(src))

	nm := copy(dest, src) //copy duplicates src into src and returns the number of items copied
	fmt.Println(src, dest, nm)

	//slice expressions
	arr := [5]int{1, 2, 3, 4, 5}

	arr1 := arr[1:3] //slicing an array returns a slice containing elements between start to stop-1 of arr
	fmt.Printf("%v, %T\n", arr1, arr1)

	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[1:3]
	fmt.Println(s2)

	fmt.Println(s1[2:]) //prints from start index 2 to the end of the slice - same as s1[2:len(s1)]
	fmt.Println(s1[:3]) //prints from index 0 to the stop index -1 - same as s1[0:3]
	fmt.Println(s1[:])  //prints all elements in the slice - same as s1[0:len(s1)]

	s1 = append(s1[:4], 100) //adds 100 on index 4
	fmt.Println(s1)

	s1 = append(s1[:4], 200) //overwrites index 4
	fmt.Println(s1)
}
