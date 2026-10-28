package main

import (
	"fmt"
)

func Multi() {
	var panjang, lebar int = 10, 5
	var (
		namaDepan    = "Budi"
		namaBelakang = "Santoso"
		akrif        = true
	)

	const (
		pi       = 3.14
		diameter = 15
	)
	var total2 float64

	total := 0
	total += panjang * lebar
	total2 += pi * diameter
	fmt.Println(namaDepan, namaBelakang, akrif, total)
}
