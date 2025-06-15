package main

import "fmt"

func main() {
	var a int = 10

	if a > 5 {
		goto greater
	}

	fmt.Println("This will be skipped")

greater:
	fmt.Println("a is greater than 5") // Output: a is greater than 5
}
