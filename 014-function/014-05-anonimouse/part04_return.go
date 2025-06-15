package main

import "fmt"

func main() {
	add := func(a, b int) int {
		return a + b
	}
	result := add(3, 4)
	fmt.Println("Result:", result) // Output: Result: 7
}
