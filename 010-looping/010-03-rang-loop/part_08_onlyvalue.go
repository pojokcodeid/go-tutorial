package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	for _, value := range numbers {
		fmt.Println("Value:", value)
	}
}
