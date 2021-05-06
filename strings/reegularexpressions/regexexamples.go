package main

import (
	"fmt"
	"regexp"
)

func main() {
	// uncompiled regex expression
	toCompare := "watermelon"
	match, _ := regexp.MatchString("w([a-z]+)melon", toCompare)
	fmt.Println("Did it match? ", match)

	// compiled match
	rCompiled, _ := regexp.Compile("w([a-z]+)melon")
	fmt.Println("Did it compile match? ", rCompiled.MatchString(toCompare))

	sentenceToCompare := "There is a watermelon in my farm. I love watermelons."
	fmt.Println("The sentence to start is: ", sentenceToCompare)
	fmt.Println("FindASubstring Results:", rCompiled.FindStringSubmatch(sentenceToCompare))
	fmt.Println("FindAllMatch Results: ", rCompiled.FindAllString(sentenceToCompare, -1))
}
