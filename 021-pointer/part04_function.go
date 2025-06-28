package main

import "fmt"

func updateValue(ptr *int) {
	*ptr = 100
}

func main() {
	var a int = 42
	fmt.Println("Before:", a) // Output: Before: 42

	updateValue(&a)
	fmt.Println("After:", a) // Output: After: 100
}
