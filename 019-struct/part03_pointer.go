package main

import "fmt"

// Defining a struct
type Person struct {
	Name string
	Age  int
}

func main() {
	// Creating a pointer to a struct
	p := &Person{Name: "Dave", Age: 35}
	fmt.Println(p)      // Output: &{Dave 35}
	fmt.Println(*p)     // Output: {Dave 35}
	fmt.Println(p.Name) // Output: Dave
	fmt.Println(p.Age)  // Output: 35
}
