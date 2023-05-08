package main

import "fmt"
import f "fmt" //imported as alias, permitted

const done = false //package scoped
var b int = 10

func main() {
	var task = "Running" //local or block scoped
	fmt.Println(task, done)
	const done = true //local scoped
	f.Printf("Done in main() is %v\n", done)
	f1()
	f.Println("Aliased FMT package print Bye!")
}

func f1() {
	const done = true
	fmt.Printf("Done in f1 is %v\n", done) //done from package scope
	a := 10
	_ = a
}
