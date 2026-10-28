package main

import (
	"fmt"
)

func KeywordMake() {
	// Slice: make(tipe, panjang) atau make(tipe, panjang, kapasitas)
	angka := make([]int, 3)     // len=3, cap=3, isinya 0 0 0
	kapas := make([]int, 3, 10) // len=3, cap=10, siap di-append

	// Map: langsung siap ditulis
	umur := make(map[string]int)
	umur["Andi"] = 27

	// Channel: dengan buffer berkapasitas 2
	ch := make(chan int, 2)
	ch <- 10

	angka[0] = 100
	fmt.Println(angka, kapas, umur["Andi"], <-ch)
}
