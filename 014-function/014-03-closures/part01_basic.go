package main

import "fmt"

// fungsi luar
func counter() func() int {
	count := 0
	// fungsi dalam (closure)
	return func() int {
		count++
		return count
	}
}

func main() {
	c := counter()

	fmt.Println(c()) // 1
	fmt.Println(c()) // 2
	fmt.Println(c()) // 3
}
