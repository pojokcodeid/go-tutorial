package main

import "fmt"

func counter2() (func() int, func() int) {
	count := 0
	increment := func() int {
		count++
		return count
	}
	decrement := func() int {
		count--
		return count
	}
	return increment, decrement
}

func main() {
	inc, dec := counter2()
	fmt.Println(inc()) // Output: 1
	fmt.Println(inc()) // Output: 2
	fmt.Println(dec()) // Output: 1
}
