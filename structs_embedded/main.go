package main

import "fmt"

func main() {

	//nested/embedded struct
	type Contact struct {
		email, address string
		phone          int
	}

	type Employee struct {
		name        string
		salary      int
		contactInfo Contact //embedded struct
	}

	john := Employee{
		name:   "John Keller",
		salary: 4000,
		contactInfo: Contact{
			email:   "jlelleer@company.com",
			address: "steetr 20, London",
			phone:   807663372,
		},
	}

	fmt.Printf("%+v\n", john)
	//accessing embedded field
	fmt.Printf("Employee's email: %s\n", john.contactInfo.email)
	//updating embedded field
	john.contactInfo.email = "new_email@company.com"
	fmt.Printf("Employee's new email: %s\n", john.contactInfo.email)
}
