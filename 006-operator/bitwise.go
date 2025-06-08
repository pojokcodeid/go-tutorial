package main

import (
	"fmt"
)

func main() {
	var a uint = 5 // 0101 in binary
	var b uint = 3 // 0011 in binary
	var result1 uint = a & b
	fmt.Println("a & b:", result1) // Output: a & b: 1 (0001 in binary)
	var result2 uint = a | b
	fmt.Println("a | b:", result2) // Output: a | b: 7 (0111 in binary)
	var result3 uint = a ^ b
	fmt.Println("a ^ b:", result3) // Output: a ^ b: 6 (0110 in binary)
	var result4 uint = a &^ b
	fmt.Println("a &^ b:", result4) // Output: a &^ b: 4 (0100 in binary)
	var result5 uint = a << 1
	fmt.Println("a << 1:", result5) // Output: a << 1: 10 (1010 in binary)
	var result6 uint = a >> 1
	fmt.Println("a >> 1:", result6) // Output: a >> 1: 2 (0010 in binary)
	var result uint = (a & b) | (a ^ b)
	fmt.Println("(a & b) | (a ^ b):", result) // Output: (a & b) | (a ^ b): 7 (0111 in binary)
}
