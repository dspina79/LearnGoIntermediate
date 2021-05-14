package main

import (
	"fmt"
	"strconv"
)

func main() {
	pl := fmt.Println

	pl(strconv.ParseInt("1929", 0, 32))
	pl(strconv.ParseInt("2134", 0, 32))
	pl(strconv.ParseInt("0xccc", 0, 32))
	result, _ := strconv.Atoi("1374")
	pl(result)
}
