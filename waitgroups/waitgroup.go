package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d starting...\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d DONE!\n", id)
}

func otherWorker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Calculating the square of %d\n", id)
	time.Sleep(1500 * time.Millisecond)
	fmt.Printf("The square of %d is %d\n", id, id*id)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
		go otherWorker(i, &wg)
	}

	wg.Wait()
}
