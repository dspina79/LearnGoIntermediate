package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	XMLName   xml.Name
	Id        int
	FirstName string
	LastName  string
	Hobbies   []string
}

func (p Person) toString() string {
	return p.FirstName + " " + p.LastName
}

func main() {
	p := Person{Id: 12, FirstName: "Dean", LastName: "Anips"}
	p.Hobbies = []string{"reading", "writing"}
	fmt.Println(p.toString())
	data, err := xml.MarshalIndent(p, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

}
