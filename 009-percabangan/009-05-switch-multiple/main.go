package main

import "fmt"

func main() {
	age := 25

	switch {
	case age < 18:
		fmt.Println("Minor")
	case age >= 18 && age <= 65:
		fmt.Println("Adult") // Output: Adult
	default:
		fmt.Println("Senior")
	}
}
