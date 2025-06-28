package main

import "fmt"

func main() {
	str := "Hello, ??"
	runes := []rune(str)

	for i, r := range runes {
		fmt.Printf("Index: %d, Rune: %c, Unicode: %U, Code point: %d\n", i, r, r, r)
	}
}
