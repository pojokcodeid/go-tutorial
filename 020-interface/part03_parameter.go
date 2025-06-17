package main

import "fmt"

// Defining an interface
type Speaker3 interface {
	Speak() string
}

type Person3 struct {
	Name string
}

func (p Person3) Speak() string {
	return "Hello, my name is " + p.Name
}

type Dog3 struct {
	Name string
}

func (d Dog3) Speak() string {
	return "Woof! I am " + d.Name
}

// Function that accepts any Speaker
func Greet(s Speaker3) {
	fmt.Println(s.Speak())
}

func main() {
	p := Person3{Name: "Charlie"}
	d := Dog3{Name: "Buddy"}

	Greet(p) // Output: Hello, my name is Charlie
	Greet(d) // Output: Woof! I am Buddy
}
