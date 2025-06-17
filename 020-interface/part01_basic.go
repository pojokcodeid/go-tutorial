package main

import "fmt"

// Defining an interface
type Speaker interface {
	Speak() string
}

// Implementing the interface in a struct
type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}

func main() {
	var s Speaker
	p := Person{Name: "Alice"}
	s = p
	fmt.Println(s.Speak()) // Output: Hello, my name is Alice
}
