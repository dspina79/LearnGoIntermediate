package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	pl := fmt.Println
	pl("Random Numbers 1-100")
	pl(rand.Intn(100))
	pl(rand.Intn(100))

	pl("\nRandom Floats")
	pl(rand.Float64())
	pl(rand.Float64())

	pl("\nNew Source")
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	pl(r1.Intn(100))
	pl(r1.Intn(100))
	pl(r1.Float64())
	pl(r1.Float64())

}
