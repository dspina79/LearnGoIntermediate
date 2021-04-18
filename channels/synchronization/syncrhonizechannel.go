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

func main() {
	exchan := make(chan bool, 1)
	go executor(exchan)

	<-exchan
}
