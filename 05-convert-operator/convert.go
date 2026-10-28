package main

import (
	"fmt"
)

func Convert() {
	var bulat int = 42
	var pecahan float64 = float64(bulat)

	var desimal float64 = 3.14
	var lagi int = int(desimal)

	fmt.Println(pecahan)
	fmt.Println(lagi)

	angka := 7
	text := fmt.Sprintf("%d", angka)
	fmt.Println("Text:", text)

	var nilai int
	fmt.Sscanf("123", "%d", &nilai)
	fmt.Println("nilai:", nilai)
}
