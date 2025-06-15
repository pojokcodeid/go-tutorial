package main

import "fmt"

func apply(operation func(int, int) int, a, b int) int {
	return operation(a, b)
}

func main() {
	result := apply(func(x, y int) int {
		return x * y
	}, 3, 4)
	fmt.Println("Multiplication Result:", result) // Output: Multiplication Result: 12
}
