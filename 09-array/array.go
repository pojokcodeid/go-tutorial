package main

import (
	"fmt"
)

func ArrayTetap() {
	// deklarasi lalu isi satu-satu
	var nilai [3]int
	nilai[0] = 10
	nilai[1] = 20
	nilai[2] = 30

	// langsung berisi nilai
	warna := [3]string{"merah", "kuning", "hijau"}

	// Panjang disimpan dari jumlah elemen dengan {...}
	prima := [...]int{2, 3, 5, 7, 11}
	fmt.Println(nilai)
	fmt.Println(warna[1])
	fmt.Println(len(prima))
}

func Salinan() {
	a := [3]int{1, 2, 3}
	b := a // b adalah SALINAN bukan rujukan

	b[0] = 90

	fmt.Println("a =", a) // [1,2,3] tidak berubah
	fmt.Println("b =", b) // [99 2 3]

	// panjang array adalah bagian dari tipe:
	var c [3]int = a // OK
	// var d [4]int = a // error type berbeda
	_ = c
}

func MultiDimensi() {
	// Array 2 dimensi : 2 baris x 3 kolom
	papan := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	for i, baris := range papan {
		for j, sel := range baris {
			fmt.Printf("papan [%d][%d] = %d \n", i, j, sel)
		}
	}
}
