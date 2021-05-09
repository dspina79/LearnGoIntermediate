package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Hobbies   []string
}

func main() {
	p := Person{FirstName: "Dean", LastName: "Anips", Age: 39}
	p.Hobbies = []string{"baseball", "reading"}

	js1, _ := json.Marshal(p)
	fmt.Println(string(js1))
}
