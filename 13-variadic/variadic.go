package main

import (
	"fmt"
)

func jumlah(angka ...int) int {
	total := 0
	for _, value := range angka {
		total += value
	}
	return total
}

func sapaSemua(sapaan string, nama ...string) {
	for _, value := range nama {
		fmt.Println(sapaan, value)
	}
}

func ContohVariadic() {
	fmt.Println(jumlah(1, 2, 3))
	fmt.Println(jumlah())
	fmt.Println(jumlah(5, 5, 5, 5, 5, 5, 5))

	sapaSemua("Halo", "Joko", "Partomo", "Partomo")

	data := []int{10, 20, 30}
	fmt.Println(jumlah(data...))
}

func Anonim() {
	// fungsi anonim disimpan pada variable
	sapa := func(nama string) string {
		return "Hallo, " + nama
	}
	fmt.Println(sapa("Joko"))

	// function langsung dipanggil IIFE
	func() {
		fmt.Println("Function Langsung dipanggil")
	}()

	// function dikirimkan sebagai arguments
	jalankan(3, 4, func(a, b int) int {
		return a * b
	})
}

func jalankan(a, b int, operasi func(int, int) int) {
	fmt.Println("Hasil kali: ", operasi(a, b))
}

// hitung() mengembalikan fungsi yang mengingat variabel "n".
func hitung() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func Closure() {
	berikutnya := hitung()
	fmt.Println(berikutnya())
	fmt.Println(berikutnya())
	fmt.Println(berikutnya())
	fmt.Println(berikutnya())
}

// variadic + closure untuk lebih flexible
func penjumlahan2(awal int) func(...int) int {
	total := awal
	return func(angka ...int) int {
		for _, n := range angka {
			total += n
		}
		return total
	}
}

func Flexible() {
	p := penjumlahan2(100)
	fmt.Println(p(1, 2, 3)) //106
	fmt.Println(p(10))      // 106 state tersimpan
}
