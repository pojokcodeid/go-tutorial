package main

import "fmt"

// Defining a struct
type Person4 struct {
	Name string
	Age  int
}

// Defining a method on the struct
func (p Person4) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
	p := Person4{Name: "Grace", Age: 50}
	p.Greet()
	// Output: Hello, my name is Grace and I am 50 years old.
}
