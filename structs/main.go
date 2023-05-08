package main

import "fmt"

func main() {
	title1, author1, year1 := "The divine Comedy", "Dante Aligheri", 1320
	title2, author2, year2 := "Macbeth", "William Shakespeare", 1620

	fmt.Println("Book1: ", title1, author1, year1)
	fmt.Println("Book2: ", title2, author2, year2)

	//defining a struct tyoe
	type book struct {
		title  string
		author string
		year   int
	}

	type book1 struct {
		title, author string //both variable of the same type can be written on the same line
		year          int
	}

	//create new variable of type book struct
	myBook := book{"The divine Comedy", "Dante Aligheri", 1320}
	fmt.Println(myBook)

	//initializing a struct with the field names
	bestBook := book{title: "Animal Farm", year: 1945, author: "George Orwell"}
	fmt.Println(bestBook)

	//retrieve and update struct fields
	lastBook := book{title: "Anna Karenina"}
	fmt.Println(lastBook.title)
	fmt.Printf("%#v\n", lastBook)
	lastBook.author = "Leo Tolstoy"
	lastBook.year = 1878
	fmt.Printf("%+v\n", lastBook) //print out both field names and values

	//comparing struct variables
	aBook := book{title: "Anna Karenina", author: "Leo Tolstoy", year: 1878}
	fmt.Println(aBook == lastBook)

	//create copy of a struct
	myBoook := aBook //we can clone a struct using the = operator and it will not reference each other unlike slices and maps
	_ = myBoook

	//Anonymous struct and anonymous struct fields
	//Anonymous struct
	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Diana",
		lastName:  "Muller",
		age:       30,
	}
	fmt.Printf("%#v\n", diana)
	fmt.Printf("Diana's age %d\n", diana.age)

	//anonymous fields
	type Book struct {
		string
		float64
		bool
	}
	b1 := Book{"1984 by George Orwell", 10.2, false}
	fmt.Printf("%#v\n", b1)

	fmt.Println(b1.string)

	type Employee struct {
		name   string
		salary int
		bool
	}

	e := Employee{"John", 40000, false}
	fmt.Printf("%#v\n", e)
	e.bool = true
	fmt.Printf("%#v\n", e)

}
