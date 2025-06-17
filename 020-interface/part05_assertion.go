package main

import "fmt"

func main() {
	var val interface{} = "Hello, Go!"

	str, ok := val.(string)
	if ok {
		fmt.Println("String value:", str) // Output: String value: Hello, Go!
	} else {
		fmt.Println("Not a string")
	}

	num, ok := val.(int)
	if ok {
		fmt.Println("Integer value:", num)
	} else {
		fmt.Println("Not an integer") // Output: Not an integer
	}
}
