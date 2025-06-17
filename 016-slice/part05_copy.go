package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))
	copy(slice2, slice1)
	fmt.Println(slice2) // Output: [1 2 3 4 5]
}
