package main

import (
	"fmt"

	"20-package-custom/matematika"
	mtk "20-package-custom/matematika" // contoh import dengan alias
)

func main() {
	fmt.Println()
	hasil := matematika.Tambah(20, 10)
	fmt.Println("penjumlahan :", hasil)

	// matematika.kali(1,3) error karena kali tidak di export dengan di tandai
	// diawali huruf kecil
	fmt.Println("KaliDua(3):", matematika.KaliDua(3))
	kali2 := mtk.KaliDua(10)
	fmt.Println(kali2)
}
