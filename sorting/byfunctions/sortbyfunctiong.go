package main

import (
	"fmt"
	"sort"
)

type sArr []string

func (s sArr) Len() int {
	return len(s)
}

func (s sArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sArr) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	names := []string{"Dave", "Larry", "Michelle", "Bo", "Fortinbras", "Leu", "Lou", "Liu"}
	sort.Sort(sArr(names))
	fmt.Println("Sorted: ", names)
}
