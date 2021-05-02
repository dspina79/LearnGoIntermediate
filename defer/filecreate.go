package main

import (
	"fmt"
	"os"
)

func createFile(fileName string) *os.File {
	fmt.Println("Creating new file", fileName)
	wd, wderr := os.Getwd()

	if wderr != nil {
		panic(wderr)
	}

	f, err := os.Create(wd + "/defer/" + fileName)

	if err != nil {
		panic(err)
	}

	return f
}

func closeFile(f *os.File) {
	fmt.Println("Closing", f.Name())
	err := f.Close()

	if err != nil {
		fmt.Println("Error closing the file: ", err.Error())
		os.Exit(1)
	}
}

func writeData(f *os.File, data string) {
	fmt.Println("Writing new information")
	fmt.Fprintln(f, data)
}

func main() {
	data := "the quick brown fox jumps over the lazy dog"
	data2 := "the lazy dog jumps over the quick brown fox"

	f := createFile("story.txt")
	defer closeFile(f)
	writeData(f, data)
	writeData(f, data2)
}
