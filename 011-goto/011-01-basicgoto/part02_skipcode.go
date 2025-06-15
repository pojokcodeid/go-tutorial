package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			goto skip
		}
		fmt.Println(i) // Output: 0 1 2 3 4
	}

skip:
	fmt.Println("Skipped loop when i was 5") // Output: Skipped loop when i was 5
}
