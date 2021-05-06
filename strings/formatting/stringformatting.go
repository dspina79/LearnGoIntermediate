package main

import (
	"fmt"
	"os"
)

type point struct {
	x int
	y int
}

func main() {

	p := point{x: 1, y: 2}
	fmt.Println("Basic Print", p)
	fmt.Printf("v %v\n", p)
	fmt.Printf("+v %+v\n", p)
	fmt.Printf("#v %#v\n", p)

	fmt.Printf("T %T\n", p)
	fmt.Printf("t %t\n", true)

	fmt.Printf("d (digit) %d\n", 123)
	fmt.Printf("b (byte) %b\n", 123)
	fmt.Printf("f (float) %f\n", 123.23)
	fmt.Printf("x (hex) %x\n", 123)

	fmt.Printf("c (char) %c\n", 68)

	fmt.Printf("e %e\n", 12300000.00)
	fmt.Printf("E %E\n", 12300000.00)
	name := "Dean"
	fmt.Printf("Hello, %s\n", name)
	fmt.Fprintf(os.Stderr, "as %s\n", "error")

}
