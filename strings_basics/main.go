package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	//strings are declared between double quotes
	//a string is a slice of bytes in Go
	//strings are utf-8 by default

	s1 := "Hi there, Go!"
	fmt.Println(s1)

	fmt.Println("He says: \"Hello!\"")
	fmt.Println(`He says: "Hello!"`) //no need to escape if using backticks

	s2 := `I like Go!` //raw string
	fmt.Println(s2)
	fmt.Println("Price: 100\nBrand: Nike")
	fmt.Println(` 
Price: 100
Brand: Nike
		`) //same as above
	fmt.Println(`C:\Users\Andrei`)
	fmt.Println("C:\\Users\\Andrei") //these two lines are equal

	//concatenating strings
	var s3 = "I love " + "Go " + "programming"
	fmt.Println(s3 + "!")

	fmt.Println("Element at index 0:", s3[0]) //prints out the dec representation of I

	//s3[5] = 'x' error a string if immutable

	fmt.Printf("%s\n", s3)
	fmt.Printf("%q\n", s3)

	//runes and bytes
	var1, var2 := 'a', 'b' //rune type

	fmt.Printf("Type: %T\tValue:%x\n", var1, var2)

	str := "tarˇÅ"
	fmt.Println(len(str))

	//string length
	name := "Golang"
	n := utf8.RuneCountInString(str)  //returns the actual length in characters(runes) of the string
	m := utf8.RuneCountInString(name) //returns the actual length of the string
	fmt.Println(n)
	fmt.Println(m)

	//slicing strings (substrings)
	str1 := "I love Golang!"
	fmt.Println(str1[2:5]) //METHOD 1: returns a slice of byte

	rs := []rune(s2) //METHOD 2 create a rune slice from the string
	fmt.Printf("%T\n", rs)
	fmt.Println(string(rs[0:3])) //create and convert a slice to string

	//manipulating string using the strings package
	p := fmt.Println //declaring an alias to the fmt.Println method
	result := strings.Contains("I love Go Programming!", "lovex")
	p(result)

	result = strings.ContainsAny("success", "xys") //checks if the string contains at least one char in the second parameter
	p(result)

	result = strings.ContainsRune("golang", 'g') //checks if a rune(char) is present in a string
	p(result)

	nn := strings.Count("cheese", "e") //counts occurrence of the substring
	p(nn)

	n = strings.Count("Five", "")
	p(n)

	p(strings.ToLower("GO PyTHON jaVA")) //prints lower case string
	p(strings.ToUpper("GO PyTHON jaVA")) //prints upper case string

	//comparing strings
	p("go" == "go") //efficient but case-sensitive
	p("GO" == "go")
	p(strings.ToLower("GO") == strings.ToLower("go")) //not efficient
	p(strings.EqualFold("GO", "go"))                  //most efficient //case-insensitive

	myStr := strings.Repeat("ab", 10) //repeats string specified number of times
	p(myStr)

	myStr = strings.Replace("192.168.0.1", ".", ":", 2) //replaces the first two . characters with :
	p(myStr)

	myStr = strings.Replace("192.168.0.1", ".", ":", -1) //replaces all . characters with :
	p(myStr)

	myStr = strings.ReplaceAll("192.168.0.1", ".", ":") //same as above, replaces all . characters with :
	p(myStr)

	ss := strings.Split("a,b,c", ",") //splits the string using , separator

	fmt.Printf("%T\n", ss)
	fmt.Printf("%#v\n", ss)

	ss = strings.Split("go got go!", "") //splits the string into runes if empty separator
	fmt.Printf("%#v\n", ss)

	ss = []string{"I", "learn", "Golang"}
	myStr = strings.Join(ss, "-") //joins string using - separator
	p(myStr)

	myStr = "Orange Green \n Blue Yellow"
	fields := strings.Fields(myStr) //split string into slices by whitespaces and newline
	fmt.Printf("%T\n", fields)
	fmt.Printf("%#v\n", fields)

	ss1 := strings.TrimSpace("\t Goodbye Windows, Welcome Linux!  ") //trims leading and trailing whitespaces
	fmt.Printf("%q\n", ss1)

	ss2 := strings.Trim("...Hello Gophers!!!!?", ".!?") //trims all specified characters from string
	fmt.Printf("%q\n", ss2)

}
