package main

import (
	"fmt"
	"time"
)

// handling calls asynchronously

func printNum(n int, callType string) {
	fmt.Printf("Calling %s n = %d\n", callType, n)
}

func main() {
	for i := 0; i < 40; i++ {
		if i%2 == 0 {
			printNum(i, "synchronous")
		} else {
			go printNum(i, "asynchronous")
		}
	}

	time.Sleep(time.Second)
}
