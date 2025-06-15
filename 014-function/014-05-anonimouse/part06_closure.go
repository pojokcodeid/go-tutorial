package main

import "fmt"

func main() {
	x := 10
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment()) // Output: 11
	fmt.Println(increment()) // Output: 12
}
