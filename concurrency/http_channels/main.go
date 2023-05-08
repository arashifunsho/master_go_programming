package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
)

// modifying to use channels
func checkAndSaveBody(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		s := fmt.Sprintf("%s is DOWN!\n", url)
		s += fmt.Sprintf("Error:%v\n", err)
		c <- s //sending into the channel

	} else {
		defer resp.Body.Close()
		s := fmt.Sprintf("%s -> Status Code: %d\n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			file := strings.Split(url, "//")[1] //taking the domain name
			file += ".txt"

			s += fmt.Sprintf("Writing response body to %s\n", file)
			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				s += fmt.Sprintf("%s \n", err)
				c <- s
			}
			s += fmt.Sprintf("%s is UP\n", url)
			c <- s
		}
	}
}
func main() {
	urls := []string{"https://golang.org", "https://www.google.com", "https://www.medium.com", "https://www.montana.edu"}

	//refactoring to use channels for concurrency

	//1. declare the channel
	c := make(chan string)

	for _, url := range urls {
		//spawn go routines with function
		go checkAndSaveBody(url, c)
		//fmt.Println(strings.Repeat("#", 20))
	}
	fmt.Println("No. of goroutines: ", runtime.NumGoroutine())

	for i := 0; i < len(urls); i++ {
		//print channel messages
		fmt.Println(<-c)
	}

}
