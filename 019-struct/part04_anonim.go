package main

import "fmt"

// Defining a struct with anonymous fields
type Person2 struct {
	string
	int
}

func main() {
	// Initializing a struct with anonymous fields
	p := Person2{"Eve", 28}
	fmt.Println(p)        // Output: {Eve 28}
	fmt.Println(p.string) // Output: Eve
	fmt.Println(p.int)    // Output: 28
}
