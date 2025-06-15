package main

import "fmt"

func main() {
	greet := func(name string) {
		fmt.Printf("Hello, %s!\n", name)
	}
	greet("Bob")     // Output: Hello, Bob!
	greet("Charlie") // Output: Hello, Charlie!
}
