package main

import (
	"fmt"
	"slices"
)

func Aritmatika() {
	a, b := 7, 3
	fmt.Println(a / b)                   // 2 (bagian pecahan di buang)
	fmt.Println(float64(a) / float64(b)) // hasilnya ada komanya

	//  penugasan majemuk
	n := 10
	n += 5
	n -= 3
	n *= 2
	n /= 4
	n++
	fmt.Println(n)
}

func Perbandingan() {
	x, y := 5, 8
	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x < y)
	fmt.Println(x >= y)

	// logika && (dan), || (atau), ! (negasi)
	umur := 20
	punyaKtp := true
	bolehMemilih := umur >= 17 && punyaKtp
	fmt.Println("Boleh memilih:", bolehMemilih)
	fmt.Println(!bolehMemilih)
}

func FungsiBawaan() {
	a, b, c := 7, 3, 10
	// dua argumen
	fmt.Println(min(a, b))
	fmt.Println(max(a, b))

	// lebih dari 2 argumen
	fmt.Println(min(a, b, c))
	fmt.Println(max(a, b, c))

	// berlaku juga untuk float64 dan string urutan alfabet
	fmt.Println(min(2.5, 1.5))
	fmt.Println(max("budi", "andi", `cindy`))

	// untuk slice menggunakan package slice, tidak bisa langsung
	nilai := []int{90, 45, 78, 60}
	fmt.Println(slices.Min(nilai), slices.Max(nilai))

	// pola praktis membatasi nilai dalam rentang
	skor := 120
	skor = min(max(skor, 0), 100)
	fmt.Println("Skor:", skor)
}

func Bitwise() {
	a, b := 6, 3
	fmt.Println(a & b)  // 2 (AND ->010)
	fmt.Println(a | b)  // 7 (OR ->111)
	fmt.Println(a ^ b)  // 5 (XOR ->101)
	fmt.Println(a << 1) // 12 (geser kiri)
	fmt.Println(a >> 1) // 3 (geser kanan)
}
