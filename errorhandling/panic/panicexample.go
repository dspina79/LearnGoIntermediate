package main

import (
	"fmt"
	"os"
)

func createFile(fn string) {
	fmt.Println("Attempting to create ", fn)
	_, err := os.Create("/tmp/" + fn)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Everything is running normally....")
	createFile("examplefile.txt")
	panic("Uh-oh, an error occurred")
}
