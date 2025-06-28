package main

import "fmt"

func main() {
	var r rune = 'a'
	fmt.Printf("Rune: %c, Unicode: %U, Code point: %d\n", r, r, r)
	// Output: Rune: a, Unicode: U+0061, Code point: 97
}
