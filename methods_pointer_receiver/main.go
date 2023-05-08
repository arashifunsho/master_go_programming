package main

import "fmt"

type car struct {
	brand string
	price int
}

func changeCar(c car, newBrand string, newPrice int) {
	c.price = newPrice
	c.brand = newBrand
}

func (c car) changeCar1(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}
func (c *car) changeCar2(newBrand string, newPrice int) {
	(*c).brand = newBrand //writing the dereferencing this way is mandatory in receiver functions
	(*c).price = newPrice
}
func main() {

	myCar := car{brand: "Audi", price: 40000}
	changeCar(myCar, "Renault", 20000)
	fmt.Println(myCar)

	myCar.changeCar1("Opel", 21000)
	fmt.Println(myCar)

	(&myCar).changeCar2("Seat", 25000) //can also use myCar.changeCar2 and the GO compiler will add the & implicitly
	fmt.Println(myCar)

	var yourCar *car
	yourCar = &myCar
	fmt.Println(*yourCar)

	//valid call ways
	yourCar.changeCar2("VW", 30000)
	fmt.Println(*yourCar)
	(*yourCar).changeCar2("Another VW", 30000)
	fmt.Println(*yourCar)
}
