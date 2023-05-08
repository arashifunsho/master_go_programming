package main

import "fmt"

func f1(n int, ch chan int) {
	ch <- n
}

func main() {

	c := make(chan int) //bi-directional channel

	//only for receiving
	c1 := make(<-chan string) //uni-directional channel
	//only for sending
	c2 := make(chan<- string) //uni-directional channel
	fmt.Printf("%T,%T,%T\n", c, c1, c2)

	go f1(10, c)
	n := <-c

	fmt.Println("Value received: ", n)
	fmt.Println("Exiting main...")
}
