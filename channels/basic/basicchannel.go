package main

import "fmt"

func main() {
	strChannel := make(chan string)

	for i := 0; i < 5; i++ {
		go func() { strChannel <- ("sample message " + fmt.Sprint(i)) }()

		channelReceiver := <-strChannel
		fmt.Println(channelReceiver)
	}
}
