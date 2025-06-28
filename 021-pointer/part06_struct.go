package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func updateAge(p *Person, newAge int) {
	p.Age = newAge
}

func main() {
	p := Person{Name: "Alice", Age: 25}
	fmt.Println("Before:", p) // Output: Before: {Alice 25}

	updateAge(&p, 30)
	fmt.Println("After:", p) // Output: After: {Alice 30}
}
