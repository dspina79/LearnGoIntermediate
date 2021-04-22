package main

import (
	"fmt"
	"time"
)

func addMessage(c chan<- string, msg string, delay float64) {
	time.Sleep(time.Second * time.Duration(delay))
	c <- msg
}

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go addMessage(c1, "first message", 2)

	select {
	case res1 := <-c1:
		fmt.Println(res1)
	case <-time.After(time.Second * 3):
		fmt.Println("Error: Timeout")
	}

	go addMessage(c2, "second message", 4)

	select {
	case res2 := <-c2:
		fmt.Println(res2)
	case <-time.After(time.Second * 3):
		fmt.Println("Error: Timeout")
	}

	// example in blocking
	go addMessage(c1, "third message", 2)
	select {

	case res3 := <-c1:
		fmt.Println("Third message received: ", res3)
	default:
		fmt.Println("Nothing to record. No activity.")
	}

	// closing

	close(c1)
	close(c2)

	fmt.Println("Jobs closed...")

}
