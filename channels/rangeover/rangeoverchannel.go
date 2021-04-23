package main

import "fmt"

func main() {
	ch := make(chan string, 4)

	ch <- "hello"
	ch <- "world"
	close(ch) // without this, the system may try to iterate over empty items or cause a deadlock

	for item := range ch {
		fmt.Println(item)
	}

}
