package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	// .C is the blocking timer channel
	<-timer1.C
	fmt.Println("First fire...")
	timer1.Reset(time.Second)

	for i := 10; i > 0; i-- {
		<-timer1.C
		fmt.Println(i)
		timer1.Reset(time.Second)
	}

	timer1Stop := timer1.Stop()
	if timer1Stop {
		fmt.Println("The timer has been stopped...")
	}
}
