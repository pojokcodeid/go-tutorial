package main

import "fmt"

func main() {
	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday") // Output: Wednesday
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Weekend")
	}
}
