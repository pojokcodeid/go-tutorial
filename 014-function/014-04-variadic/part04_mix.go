package main

import "fmt"

func printAll(values ...interface{}) {
	for _, value := range values {
		fmt.Println(value)
	}
}

func main() {
	printAll(1, "hello", 3.14, true)
	// Output:
	// 1
	// hello
	// 3.14
	// true
}
