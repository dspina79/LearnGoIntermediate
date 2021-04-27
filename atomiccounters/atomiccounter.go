package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter1 uint64
	var counter2 uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&counter1, 1)
				if counter1%2 == 0 {
					atomic.AddUint64(&counter2, 1)
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter1 Count", counter1)
	fmt.Println("Counter2 Count (Likely Blocked)", counter2)
}
