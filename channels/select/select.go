package main

import (
	"fmt"
	"time"
)

func calAreaWithDelay(c chan<- int, side int) {
	time.Sleep(time.Second * 3)
	c <- (side * side)
}

func generateGreetingWithDelay(c chan<- string, msg string) {
	time.Sleep(time.Second * 7)
	c <- ("Hello " + msg)
}

func main() {
	intChan := make(chan int, 1)
	stringChan := make(chan string, 1)
	x := true
	y := true
	go calAreaWithDelay(intChan, 4)
	go generateGreetingWithDelay(stringChan, "person")

	fmt.Println("Entering select...")

	for x || y {
		select {
		case resp1 := <-stringChan:
			fmt.Println("received string: ", resp1)
			x = false
		case resp2 := <-intChan:
			fmt.Println("received int: ", resp2)
			y = false
		}
	}

	fmt.Println("At the end of the requests...")
}
