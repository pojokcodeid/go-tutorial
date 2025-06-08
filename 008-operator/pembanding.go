package main

import (
	"fmt"
)

func main() {
	value := (((2+6)%3)*4 - 2) / 3
	isEqual := (value == 2)

	fmt.Printf("nilai %d (%t) \n", value, isEqual)
}
