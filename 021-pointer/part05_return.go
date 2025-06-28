package main

import "fmt"

func createPointer() *int {
	var a int = 42
	return &a
}

func main() {
	ptr := createPointer()
	fmt.Println(*ptr) // Output: 42
}
