package main

import (
	"fmt"
)

func ContohPonter() {
	x := 20

	// p menyimpan alamat dari x
	var p *int = &x

	fmt.Println("Nilai x:", x)
	fmt.Println("Alamat x:", p)
	fmt.Println("Isi alamat x:", *p)

	// mengubah alamat x lewat pointer
	*p = 45
	fmt.Println("x sekarang :", x)
}

// menerima salinan tidak merubah nilai asli
func gandakanSalah(n int) {
	n = n * 2
}

// menerima pointer dan mengubah nilai asli
func gandakanBenar(n *int) {
	*n = *n * 2
}

func PassByPointer() {
	angka := 10
	gandakanSalah(angka)
	fmt.Println("Contoh gandakan salah:", angka)

	gandakanBenar(&angka)
	fmt.Println("Contoh gandakan beanr :", angka)
}

type Titik struct {
	X, Y int
}

// pinter ke struct bisa di ubah field-nya tanpa di reference explicit
func geser(t *Titik, dx, dy int) {
	t.X += dx // otomatis dianggap *t.X
	t.Y += dy
}

func PointerNil() {
	var p *int
	fmt.Println(p)
	if p == nil { // akan ada warning linter, bisa diabaikan karena ini untuk membuktikans aja
		fmt.Println("Pointer belum merujuk ke manapun")
	}

	titik := &Titik{X: 1, Y: 2}
	geser(titik, 10, 2)
	fmt.Println("nilai titik:", *titik)
}
