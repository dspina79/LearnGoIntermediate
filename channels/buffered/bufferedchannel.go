package main

import "fmt"

func main() {
	c := make(chan string, 3)

	c <- "item 1"
	c <- "item 2"
	c <- "item 3"

	for i := 0; i < 3; i++ {
		fmt.Println(<-c)
	}

}
