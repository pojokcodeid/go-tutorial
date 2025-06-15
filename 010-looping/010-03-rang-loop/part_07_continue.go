package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // Skip the rest of the loop for even numbers
		}
		fmt.Println(i) // Output: 1 3 5 7 9
	}
}
