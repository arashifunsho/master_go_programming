package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	const gr = 100

	var wg sync.WaitGroup
	wg.Add(gr * 2)

	var n int = 0
	//using mutex lock to prevent data races
	//1 include the mutex variable
	var m sync.Mutex

	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			//2. block access to n until this goroutine finishes
			m.Lock()
			n++
			m.Unlock()
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			defer m.Unlock() //also works
			n--
			wg.Done()
		}()
	}

	fmt.Println("No of GoRoutines: ", runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("Final Value of n: ", n)
}
