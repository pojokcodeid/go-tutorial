package main

import "fmt"

func main() {
	numberOfDays := 28

	// switch without any expression
	switch {

	case 28 == numberOfDays:
		fmt.Println("It's February")

	default:
		fmt.Println("Not February")
	}
}
