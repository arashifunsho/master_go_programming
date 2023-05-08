package main

import "fmt"

func main() {
	var ch chan int //declaring a channel variable
	fmt.Println(ch)

	ch = make(chan int) //initializing a new channel
	fmt.Println(ch)
	c := make(chan int) //declaring and initializing a new channel

	// <- channel operator

	c <- 10 //sends the value 10 into the channel

	num := <-c //receive the value from the channel into a variable
	fmt.Println(<-c)
	_ = num

	//close a channel
	close(c)

}
