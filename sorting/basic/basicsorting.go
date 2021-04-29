package main

import (
	"fmt"
	"sort"
)

func main() {
	alphabeticalStringArray := []string{"c", "a", "k", "e"}
	numericArray := []int{3, 6, 10, 2, 2340, 92, 99, 1}
	wordArray := []string{"The", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	mixedAbcArray := []string{"z", "E", "a", "A", "Z", "e"}

	sort.Strings(alphabeticalStringArray)
	sort.Strings(wordArray)
	sort.Strings(mixedAbcArray)
	sort.Ints(numericArray)

	fmt.Println("Alphabetical Array", alphabeticalStringArray)
	fmt.Println("Word Array", wordArray)
	fmt.Println("Mixed Alphabetical Array", mixedAbcArray)

	fmt.Println("Numeric Array", numericArray)

	// therefore, A < a and The < the
}
