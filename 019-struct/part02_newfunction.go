package main

import "fmt"

// Defining a struct
type Person struct {
	Name string
	Age  int
}

func main() {
	// Initializing a struct using the new function
	p := new(Person)
	p.Name = "Charlie"
	p.Age = 40
	fmt.Println(*p) // Output: {Charlie 40}
}
