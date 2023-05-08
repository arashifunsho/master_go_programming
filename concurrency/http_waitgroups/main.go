package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

func checkAndSaveBody(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is DOWN!\n", url)
	} else {
		defer resp.Body.Close()
		fmt.Printf("%s -> Status Code: %d\n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			file := strings.Split(url, "//")[1] //taking the domain name
			file += ".txt"

			fmt.Printf("Writing response body to %s\n", file)
			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	fmt.Println(strings.Repeat("#", 20))
	wg.Done() //send done signal to waitgroup variable
}
func main() {
	urls := []string{"https://golang.org", "https://www.google.com", "https://www.medium.com", "https://www.montana.edu"}

	var wg sync.WaitGroup //init wait group for concurrent go routines
	wg.Add(len(urls))     //number of go routines to wait for

	for _, url := range urls {
		go checkAndSaveBody(url, &wg)
	}
	fmt.Println("No of Goroutines: ", runtime.NumGoroutine())

	wg.Wait() //wait for go routines to complete execution

}
