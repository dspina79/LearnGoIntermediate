package main

import (
	"encoding/xml"
	"fmt"
)

type person struct {
	XMLName   xml.Name
	id        int
	firstName string
	lastName  string
	hobbies   []string
}

func (p person) toString() string {
	return p.firstName + " " + p.lastName
}

func main() {
	p := &person{id: 12, firstName: "Dean", lastName: "Anips"}
	p.hobbies = []string{"reading", "writing"}
	fmt.Println(p.toString())
	data, _ := xml.MarshalIndent(p, " ", "  ")
	fmt.Println(string(data))

}
