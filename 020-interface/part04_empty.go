package main

import "fmt"

func PrintAnyType(val interface{}) {
	fmt.Println(val)
}

func main() {
	PrintAnyType("Hello, World!") // Output: Hello, World!
	PrintAnyType(42)              // Output: 42
	PrintAnyType(true)            // Output: true
}
