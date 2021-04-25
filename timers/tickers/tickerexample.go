package main

import (
	"fmt"
	"time"
)

func message(msg string, tick time.Ticker, sentinel chan bool) {
	for {
		select {
		case <-sentinel:
			fmt.Println("All done")
		case <-tick.C:
			fmt.Println("Ticker at ", msg)
		}
	}
}

func main() {
	tick := time.NewTicker(time.Second)
	sent := make(chan bool, 1)

	go message("Hello", *tick, sent)

	time.Sleep(10 * time.Second)
	tick.Stop()
	sent <- true
}
