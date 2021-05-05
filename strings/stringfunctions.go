package main

import (
	"fmt"
	s "strings"
)

func main() {

	fmt.Println("Some Basic String Functions")

	fmt.Println("Contains\t\t\t\t", s.Contains("the quick and the slow", "the"))
	fmt.Println("Count\t\t\t", s.Count("supression", "s"))
	fmt.Println("Upper\t\t", s.ToUpper("this was lowercase"))
	fmt.Println("Lower\t\t", s.ToLower("THIS WAS ALL UPPERCASE"))
	fmt.Println("HasPrefix/Begins\t\t", s.HasPrefix("better", "be"))
	fmt.Println("HasPrefix/Begins\t\t", s.HasPrefix("better", "te"))
	fmt.Println("HasSuffix/Ends\t\t\t", s.HasSuffix("better", "er"))
	fmt.Println("HasSuffix/Ends\t\t\t", s.HasSuffix("better", "be"))
	fmt.Println("Index\t\t", s.Index("better", "e"))
	fmt.Println("Join\t\t", s.Join([]string{"a", "b", "c"}, "|"))
	fmt.Println("Repalce\t\t", s.Replace("Dave", "av", "eb", -1))
	fmt.Println("Repeat\t\t", s.Repeat("^", 10))
	fmt.Println("Split\t\t", s.Split("abcdefghi", ""))

}
