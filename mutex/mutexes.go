package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func readData(state map[int]int, key int) int {
	return state[key]
}

func writeData(state map[int]int, key int, val int) {
	state[key] = val
}

func main() {
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}

	var readOperations uint64
	var writeOperations uint64

	for r := 0; r < 100; r++ {
		total := 0
		key := rand.Intn(5)
		mutex.Lock()
		total += readData(state, key)
		mutex.Unlock()
		atomic.AddUint64(&readOperations, 1)
		time.Sleep(time.Millisecond)
	}

	for r := 0; r < 100; r++ {
		key := rand.Intn(5)
		val := rand.Intn(100)
		mutex.Lock()
		writeData(state, key, val)
		mutex.Unlock()
		atomic.AddUint64(&writeOperations, 1)
		time.Sleep(time.Millisecond)
	}

	time.Sleep(time.Second)

	readOpsCount := atomic.LoadUint64(&readOperations)
	writeOpsCount := atomic.LoadUint64(&writeOperations)

	fmt.Println("Reads: ", readOpsCount)

	fmt.Println("Writes: ", writeOpsCount)

	mutex.Lock()
	fmt.Println("state: ", state)
	mutex.Unlock()
}
