package main

import "fmt"

func sum2(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(sum2(numbers...)) // Output: 15
}
