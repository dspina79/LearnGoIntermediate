package main

import "fmt"

func receive(incoming chan<- string, msg string) {
	incoming <- msg
}

func transfer(incoming <-chan string, outgoing chan<- string) {
	received := <-incoming
	outgoing <- received
}

func main() {
	mainChan := make(chan string, 1)
	subChan := make(chan string, 1)
	receive(mainChan, "hello word")
	transfer(mainChan, subChan)
	receivedMessage := <-subChan
	fmt.Println(receivedMessage)
	receive(subChan, "good-bye")

	fmt.Println(<-subChan)
}
