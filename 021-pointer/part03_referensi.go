package main

import "fmt"

func main() {
	var a int = 42
	var ptr *int = &a
	fmt.Println(*ptr) // Output: 42

	*ptr = 21
	fmt.Println(a) // Output: 21
}
