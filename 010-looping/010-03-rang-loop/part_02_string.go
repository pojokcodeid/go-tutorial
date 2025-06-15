package main

import "fmt"

func main() {
	str := "Hello"
	for index, char := range str {
		fmt.Printf("Index: %d, Char: %c\n", index, char)
	}
}
