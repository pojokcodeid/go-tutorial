package main

import "fmt"

// Defining an interface
type Speaker2 interface {
	Speak2() string
}

// Implementing the interface in a struct
type Person2 struct {
	Name string
}

func (p Person2) Speak2() string {
	return "Hello, my name is " + p.Name
}

type Dog struct {
	Name string
}

func (d Dog) Speak2() string {
	return "Woof! I am " + d.Name
}

func main() {
	var speakers []Speaker2

	p := Person2{Name: "Bob"}
	d := Dog{Name: "Rex"}

	speakers = append(speakers, p, d)

	for _, s := range speakers {
		fmt.Println(s.Speak2())
	}
	// Output:
	// Hello, my name is Bob
	// Woof! I am Rex
}
