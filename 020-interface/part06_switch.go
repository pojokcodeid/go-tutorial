package main

import "fmt"

func Describe(val interface{}) {
	switch v := val.(type) {
	case string:
		fmt.Println("String value:", v)
	case int:
		fmt.Println("Integer value:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	Describe("Hello, Go!") // Output: String value: Hello, Go!
	Describe(42)           // Output: Integer value: 42
	Describe(true)         // Output: Unknown type
}
