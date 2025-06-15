package main

import "fmt"

func main() {
	num := 2

	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two") // Output: Two
		fallthrough
	case 3:
		fmt.Println("Three") // Output: Three
	default:
		fmt.Println("Other")
	}
}
