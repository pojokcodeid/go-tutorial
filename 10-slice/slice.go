package main

import (
	"fmt"
)

func MembuatSlace() {
	// 1) Literal (tanpa panjang) -> slice
	buah := []string{"Apel", "Jeruk", "Mangga"}
	// 2) Dengan make:panjang dan kapasitas
	angka := make([]int, 3, 5) // len= 3 cap=5 isi=0
	// 3) Slice kosong (nil)
	var kosong []int

	fmt.Println(buah, len(buah), cap(buah))
	fmt.Println(angka, len(angka), cap(angka))
	fmt.Println(kosong, len(kosong), kosong == nil)
}

func MenambahElement() {
	s := make([]int, 0, 2)
	fmt.Println(len(s), cap(s)) // 0 2

	s = append(s, 10)
	fmt.Println(len(s), cap(s)) // 1 2

	s = append(s, 20)
	fmt.Println(len(s), cap(s)) // 2 2

	// kapasitas penuh -> Go mengalokasi array baru
	s = append(s, 30)
	fmt.Println(len(s), cap(s)) // 3 4 (cap biasanya berlipat)
	fmt.Println(s)
}

func Slicing() {
	asli := []int{0, 1, 2, 3, 4, 5}

	potong := asli[1:4]                   // index 1 sampai 3 (4 tidak termasuk)
	fmt.Println(potong)                   // [1,2,3]
	fmt.Println(len(potong), cap(potong)) // 3 5

	// potongan berbagi aray dasar dengan asli
	potong[0] = 99
	fmt.Println("Asli:", asli) // [0,99,2,3,4,5]

	// salinan sebenarnya -> pakai copy
	salinan := make([]int, len(potong))
	copy(salinan, potong)
	salinan[0] = 7
	fmt.Println(potong, salinan) // [99 2 3] [7 2 3]
}

func MenghapusSlice() {
	s := []int{10, 20, 30, 40, 50}

	// hapus element index 2 (nilai 30)
	i := 2
	s = append(s[:i], s[i+1:]...)
	fmt.Println(s) // [10,20,30,40,50]

	// menyalin variable slice juga berbagi array dasar -> berhati-hatilah
	a := []int{1, 2, 3}
	b := a
	b[0] = 100
	fmt.Println(a) // [100,2,3]
}
