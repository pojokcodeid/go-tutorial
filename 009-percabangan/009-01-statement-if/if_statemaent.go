package main

import "fmt"

func main() {
	number := 15

	// bernilai true jika number lebih besar dari 0
	if number > 0 {
		fmt.Printf("%d is a positive number\n", number)
	}

	fmt.Println("Out of the loop")
}
