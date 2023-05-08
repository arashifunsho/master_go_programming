package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// The pattern to use sync.WaitGroup is:
// 1. Create a new variables of a `sync.WaitGroup` (wg)
// 2. Call `wg.Add(n)` where `n` is the number of goroutines to wait for
// 3. Execute `defer wg.Done()` in each goroutine to indicate to the WaitGroup that the goroutine has finished executing
// 4. Call `wg.Wait()` in main() where we want to block.

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 (goroutine) execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1, i=", i)
		time.Sleep(time.Second)
	}
	fmt.Println("f1 execution finished")
	//step 3
	wg.Done() //or (*wg).Done()
}
func f2() {
	fmt.Println("f2 execution started")
	for i := 5; i < 8; i++ {
		fmt.Println("f2, i=", i)
	}
	fmt.Println("f2 execution finished")
}

func main() {
	//step 1
	var wg sync.WaitGroup //using waitgroup to wait till all goroutines has finished executing before exiting the program
	//step 2
	wg.Add(1)

	go f1(&wg) //running f1() as a goroutine i.e concurrently
	fmt.Println("No. of Goroutines after go f1():", runtime.NumGoroutine())
	f2()

	//step 4
	wg.Wait() //blocking the main function till all go routines has been executed
	fmt.Println("main execution stopped")

}
