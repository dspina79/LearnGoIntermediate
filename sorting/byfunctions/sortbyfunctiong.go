package main

import (
	"fmt"
	"sort"
	"strings"
)

type byLastChar []string
type sArr []string

func (s byLastChar) Len() int {
	return len(s)
}

func (s byLastChar) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLastChar) Less(i, j int) bool {
	si := strings.ToLower(strings.Split(s[i], "")[len(s[i])-1])
	sj := strings.ToLower(strings.Split(s[j], "")[len(s[j])-1])
	return si < sj
}

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
	names2 := []string{"Dave", "Larry", "Michelle", "Bo", "Fortinbras", "Leu", "Lou", "Liu"}
	sort.Sort(sArr(names))
	sort.Sort(byLastChar(names2))
	fmt.Println("Sorted by Length: ", names)
	fmt.Println("Sorted by Last Char: ", names2)
}
