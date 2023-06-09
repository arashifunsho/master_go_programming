package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)
	fmt.Println("Channel's capacity:", cap(c)) // => 3

	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("func goroutine #%d before sending data into the channel\n", i)
			c <- i
			fmt.Printf("func goroutine #%d after sending data into the channel\n", i)
		}
		close(c)

	}(c)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	for v := range c {
		fmt.Println("main goroutine received value from channel: ", v)
	}
	fmt.Println("Receiving from a closed channel prints the default in this case, ", <-c)

	//c <- 2 //sending on closed channel results in go panic
}
