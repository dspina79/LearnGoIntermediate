package main

import (
	"fmt"
	"math/rand"
)

func check(x int, y int, answerProvided int) {
	if (x + y) == answerProvided {
		fmt.Println("GREAT JOB!!! YOU'RE CORRECT!")
	} else {
		fmt.Printf("Sorry, the right answer to %d + %d is %d\n", x, y, (x + y))
	}
}

func main() {
	for i := 0; i < 10; i++ {
		x := rand.Intn(10)
		y := rand.Intn(10)
		var ans int

		fmt.Printf("%d + %d = ", x, y)
		fmt.Scanf("%d", &ans)
		check(x, y, ans)
	}
}
