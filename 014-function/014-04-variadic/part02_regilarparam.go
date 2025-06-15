package main

import "fmt"

func greet(prefix string, names ...string) {
	for _, name := range names {
		fmt.Printf("%s %s\n", prefix, name)
	}
}

func main() {
	greet("Hello", "Alice", "Bob", "Charlie")
	// Output:
	// Hello Alice
	// Hello Bob
	// Hello Charlie
}
