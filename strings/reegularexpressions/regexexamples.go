package main

import (
	"fmt"
	"regexp"
)

func main() {
	// uncompiled regex expression
	toCompare := "watermelon"
	match, _ := regexp.MatchString("w([a-z]+)elon", toCompare)
	fmt.Println("Did it match? ", match)

	// compiled match
	rCompiled, _ := regexp.Compile("w([a-z]+)elon")
	fmt.Println("Did it compile match? ", rCompiled.MatchString(toCompare))

	sentenceToCompare := "There is a watermelon in my farm. I love watermelons."
	fmt.Println("The sentence to start is: ", sentenceToCompare)
	fmt.Println("FindStringSubmatch Results:", rCompiled.FindStringSubmatch(sentenceToCompare))
	fmt.Println("FindAllString Results: ", rCompiled.FindAllString(sentenceToCompare, -1))
	fmt.Println("FindAllStringIndex Result:", rCompiled.FindAllStringIndex(sentenceToCompare, -1))
	fmt.Println("FindStringSubmatchIndex Result:", rCompiled.FindStringSubmatchIndex(sentenceToCompare))
}
