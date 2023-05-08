package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"b.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() //postponing closing the file till later

	//writing to file using os
	byteSlice := []byte("I learn Golang!")
	byteWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes Written: %d\n", byteWritten)

	//writing to file using ioutil.WriteFile()
	bs := []byte("Go Programming is cool!")
	err = ioutil.WriteFile("c.txt", bs, 0644)
	if err != nil {
		log.Fatal(err)
	}

	//writing to file using bufio package
	file1, err := os.OpenFile("my file.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()

	//writing to buffer
	bufferedWriter := bufio.NewWriter(file1)
	bs1 := []byte{97, 98, 99}
	byteWritten, err = bufferedWriter.Write(bs1)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Bytes written to buffer (not file) %d\n", byteWritten)

	bytesAvailable := bufferedWriter.Available()
	log.Printf("Bytes available in buffer %d\n", bytesAvailable)

	byteWritten, err = bufferedWriter.WriteString("\nJust a random string")
	if err != nil {
		log.Fatal(err)
	}
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered  %d\n", unflushedBufferSize)

	//flushing buffer to file
	err = bufferedWriter.Flush()
	if err != nil {
		log.Fatal(err)
	}

	//discard content of buffer (effective if buffer hasn't been flushed)
	bufferedWriter.Reset(bufferedWriter)
}
