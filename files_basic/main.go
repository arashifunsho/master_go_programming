package main

import (
	"fmt"
	"log"
	"os"
)

//os package is used for basic file operations
//other packages such as io, ioutil, bufio are reserved for advanced operations

func main() {
	//creating a new file
	var newFile *os.File
	fmt.Printf("%T\n", newFile)

	var err error
	newFile, err = os.Create("a.txt") //will create a.txt if it doesn't already exist or truncates it if it does
	if err != nil {
		//method one for printing out errors
		//fmt.Println(err)
		//os.Exit(1)

		//method 2 and the idiomatic way
		log.Fatal(err)

	}
	fmt.Printf("%T\n", newFile)
	err = os.Truncate("a.txt", 0) //empties the file or truncates to specified byte
	if err != nil {
		log.Fatal(err)
	}

	//closing the file
	err = newFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	//open an existing file
	//file, err := os.Open("a.txt") //default
	file, err := os.OpenFile("a.txt", os.O_CREATE|os.O_APPEND, 0644) //open in create if it does not exist or append to file if it does
	if err != nil {
		log.Fatal(err)
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	//get file information
	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("a.txt")
	p := fmt.Println

	p("File Name: ", fileInfo.Name())
	p("File Size (bytes): ", fileInfo.Size())
	p("Last Modified: ", fileInfo.ModTime())
	p("is Directory: ", fileInfo.IsDir())
	p("Permissions: ", fileInfo.Mode())
	fileInfo, err = os.Stat("b.txt")
	if err != nil {
		p(err)
		if os.IsNotExist(err) {
			p("File does not exist")
		}
	}

	//rename or move a file
	oldPath, newPath := "a.txt", "aaa.txt"

	err = os.Rename(oldPath, newPath)
	if err != nil {
		p(err)
	}

	//remove a file
	err = os.Remove(newPath)
	if err != nil {
		log.Fatal(err)
	}

}
