package main

import "fmt"

func main() {
	var a int = 42
	var ptr *int = &a
	fmt.Println(ptr) // Output: Memory address of a
}
