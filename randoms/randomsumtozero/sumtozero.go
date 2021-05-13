package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandom(n int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	test := r1.Intn(2)
	val := r1.Intn(n)
	if test == 0 {
		return val * -1
	}

	return val
}

func main() {
	iterations := 0
	sum := -101

	for sum != 0 && iterations < 1000000 {
		if sum == -101 {
			sum = 0
		}
		iterations += 1
		sum += getRandom(100)
	}

	fmt.Println("Total iterations to get to 0:", iterations)
}
