package main

import "fmt"

func main() {
	//appends is an expensive operation so internally go tries to increase(create new) the backing array as little as possible
	//appends assigns a growth factor of 2 i.e it multiplies the current backing array by 2 if it needs to create a new backing array as slices grow
	var nums []int
	fmt.Printf("%#v\n", nums)
	fmt.Printf("Length: %d\tCapacity:%d\n", len(nums), cap(nums))

	nums = append(nums, 1, 2)
	fmt.Printf("Length: %d\tCapacity:%d\n", len(nums), cap(nums))

	nums = append(nums, 3)
	fmt.Printf("Length: %d\tCapacity:%d\n", len(nums), cap(nums))

	nums = append(nums, 4)
	fmt.Printf("Length: %d\tCapacity:%d\n", len(nums), cap(nums))
	nums = append(nums, 100)
	fmt.Printf("Length: %d\tCapacity:%d\n", len(nums), cap(nums))
	nums = append(nums[0:4], 200, 300, 400, 500, 600)
	fmt.Printf("Length: %d\tCapacity:%d\n", len(nums), cap(nums))

	letters := []string{"A", "B", "C", "D", "E", "F"}
	letters = append(letters[:1], "X", "Y")
	fmt.Printf("Slice: %v\tLength: %d\tCapacity:%d\n", letters, len(letters), cap(letters))
	//fmt.Println(letters[4])//error index out of bound even though we know the capacity is 6.
	fmt.Println(letters[3:6]) //slice expression can access beyond the length because they have access to the backing array
}
