package main

import "fmt"

func square(c chan int, num int) {
	c <- (num * num)
}

func main() {
	strChannel := make(chan string)
	intChannel := make(chan int)

	go square(intChannel, 13)
	go square(intChannel, 14)

	topStack := <-intChannel
	bottomStack := <-intChannel

	fmt.Println("Top: ", topStack, " Bottom: ", bottomStack)

	for i := 0; i < 5; i++ {
		go func() { strChannel <- ("sample message " + fmt.Sprint(i)) }()

		channelReceiver := <-strChannel
		fmt.Println(channelReceiver)
	}
}
