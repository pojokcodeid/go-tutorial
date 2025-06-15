package main

import "fmt"

func main() {
	var x interface{} = 10

	switch v := x.(type) {
	case int:
		fmt.Println("x is an int:", v) // Output: x is an int
	case string:
		fmt.Println("x is a string:", v)
	default:
		fmt.Println("Unknown type")
	}
}
