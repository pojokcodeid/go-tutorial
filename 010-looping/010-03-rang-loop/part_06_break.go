package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // Exit the loop when i is 5
		}
		fmt.Println(i)
	}
}
