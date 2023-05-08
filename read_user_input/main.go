package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) //takes input from commandline
	scanner.Scan()

	text := scanner.Text()
	bytes := scanner.Bytes()

	fmt.Println("Input text: ", text)
	fmt.Println("Input bytes: ", bytes)

	//continues reading

	for scanner.Scan() {
		text = scanner.Text()
		fmt.Println("You entered: ", text)
		if text == "exit" {
			fmt.Println("Exiting the scanning")
			break
		}
	}

	fmt.Println("Just a message after the for loop")
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
