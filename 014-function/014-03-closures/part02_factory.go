package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	add := adder()
	fmt.Println(add(10)) // Output: 10
	fmt.Println(add(5))  // Output: 15
	fmt.Println(add(2))  // Output: 17
}
