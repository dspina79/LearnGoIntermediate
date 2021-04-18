package main

import (
	"fmt"
	"time"
)

func executor(c chan bool) {
	fmt.Println("Executing Channel")
	time.Sleep(time.Second)
	fmt.Println("Finished")

	c <- true // without this, the system will deadlock
}

func printNumbers(c chan bool) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Running number %d\n", i)
		time.Sleep(time.Millisecond * 500)
	}

	c <- true
}

func main() {
	exchan := make(chan bool, 1)
	numChan := make(chan bool, 1)
	go executor(exchan)
	go printNumbers(numChan)
	<-numChan
	<-exchan
}
