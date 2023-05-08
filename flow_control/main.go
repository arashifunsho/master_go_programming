package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	//If statements
	price, inStock := 100, true
	if price > 80 {
		fmt.Println("Too Expensive")
	}

	_ = inStock

	if price <= 100 && inStock {
		fmt.Println("Buy it!")
	}

	if price < 100 {
		fmt.Println("It's cheap")
	} else if price == 100 {
		fmt.Println("On the edge")
	} else {
		fmt.Println("It's expensive")
	}

	//simple if statement
	i, err := strconv.Atoi("45a")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	//equivalent of the above but simpler
	if i, err := strconv.Atoi("20"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("No error, i is:", i)
	}

	if args := os.Args; len(args) != 2 {
		fmt.Println("One argument is required!")
	} else if km, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("The argument must be an integer! Error:", err)
	} else {
		fmt.Printf("%d km in miles is %v\n", km, float64(km)/1.609)
	}

	//FOR loops
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//no while loop but this can simulate it in go
	j := 10
	for j >= 0 {
		fmt.Println(j)
		j--
	}
	//infinite loop
	//sum := 0
	//for {
	//	sum++
	//}
	//fmt.Println(sum) //line is never reached

	//CONTINUE STATEMENT
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}

		fmt.Println(i)
	}

	//BREAK STATEMENT
	count := 0
	for i := 0; true; i++ {
		if i > 0 && i%13 == 0 {
			fmt.Printf("%d is divisible by 13\n", i)
			count++
		}
		if count == 10 {
			break
		}
	}
	fmt.Println("Just a message after the for loop")

	//LABEL STATEMENT
	outer := 19 //proof that label names don't conflict with identifier names
	_ = outer
	people := [5]string{"Helen", "Mark", "Antonio", "Brenda", "Michael"}
	friends := [2]string{"Mark", "Mary"}
outer:
	for index, name := range people { //iterating over an array
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("Found a friend %q at index %d\n", friend, index)
				break outer
			}
		}
	}
	fmt.Println("Next instruction after the break")

	//GOTO STATEMENT
	v := 0
loop: //goto causes a loop to occur
	if v < 5 {
		fmt.Println(v)
		v++
		goto loop
	}
	//	goto todo1
	//	x := 5 //not allowed because a new variable came into scope
	//todo1:
	//	fmt.Println("Something here")

	//SWITCH STATEMENT
	language := "golang"
	switch language {
	case "Python":
		fmt.Println("You are learning Python!") //implicit break is added after each case
	case "Go", "golang": //multiple case that match a scope
		fmt.Println("Good Go in view")

	default:
		fmt.Println("Any other is a good start too")
	}
	n := 5
	switch { //same as switch true
	case n%2 == 0:
		fmt.Println("Even Number")
	case n%2 != 0:
		fmt.Println("Odd Number")
	}

}
