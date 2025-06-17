package main

import "fmt"

func main() {
	slice := make([]int, 3, 5)
	fmt.Println(slice)      // Output: [0 0 0]
	fmt.Println(len(slice)) // Output: 3
	fmt.Println(cap(slice)) // Output: 5
}
