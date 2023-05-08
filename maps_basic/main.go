package main

import "fmt"

func main() {
	var employees map[string]string //map with string keys and values uninitialized
	fmt.Printf("%#v\n", employees)

	fmt.Printf("number of pairs: %d\n", len(employees))
	fmt.Printf("Value for key %q is %q\n", "Dan", employees["Dan"]) //returns empty string if map at key is does not exist

	var accounts map[string]float64
	fmt.Printf("%#v\n", accounts["0x323"])

	//employees["Dan"] = "Programmer" error map has to be initialized with make function or literal before assigning value

	people := map[string]float64{} //empty but initialized map

	people["John"] = 21.4
	people["Mary"] = 10
	fmt.Println(people)

	map1 := make(map[int]int) //another method to initialize a map
	map1[4] = 8

	balances := map[string]float64{ //declare a map and initialize using literals
		"USD": 323.1,
		"EUR": 432.14,
		//50:322, all keys must be the same type as well as all values
		"CHF": 3234.6,
	}

	fmt.Println(balances)
	m := map[int]int{1: 10, 2: 20, 3: 30}
	_ = m

	//adding/updating key
	balances["USD"] = 500.2  //updates value of existing key
	balances["GBP"] = 120.87 //adds new key

	fmt.Println(balances)
	balances["RON"] = 0 //adds new key

	fmt.Println(balances["RON"]) //returns the default empty variable of the type if key does not exist

	v, exists := balances["RON"] //checking if it's a zero value or non-existing key; exists will be true if key exists
	if exists {
		fmt.Println("The RON balance is: ", v)
	} else {
		fmt.Println("The RON key does not exist in the map")
	}

	//iterating  a map
	for k, v := range balances { //doesn't maintain order
		fmt.Printf("key: %v,\tValue: %v\n", k, v)
	}

	//removing a key from a map
	delete(balances, "USD")
	fmt.Println(balances)

	//comparing maps
	a := map[string]string{"A": "X"}
	b := map[string]string{"A": "X"}

	//fmt.Println(a == b) error cannot compare maps this way

	//comparing via string representation if both key and value are strings
	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)

	fmt.Println(s1, s2)
	if s1 == s2 {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}

	//cloning or duplicating a map
	friends := map[string]int{"Dan": 40, "Maria": 25}

	neighbours := friends //neighbours and friends have the same map header and reference the same map in memory

	friends["Dan"] = 50

	fmt.Println(neighbours)

	people1 := make(map[string]int)
	for k, v := range friends { //copying from one map into another
		people1[k] = v
	}

	people1["Anne"] = 18
	fmt.Printf("%#v\t\t%#v\n", people1, friends)
}
