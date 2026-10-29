package main

import (
	"fmt"
)

func ContohSwitch1() {
	hari := 6
	switch hari {
	case 1, 2, 3, 4, 5:
		fmt.Println("Hari kerja")
	case 6, 7:
		fmt.Println("Hari libur")
	default:
		fmt.Println("Hari tidak diketahui")
	}
}

func TanpaExpresi() {
	// Switch tanpa ekspresi: setiap case adalah kondisi bool
	nilai := 78
	switch {
	case nilai >= 85:
		fmt.Println("A")
	case nilai >= 70:
		fmt.Println("B")
	default:
		fmt.Println("C atau Kurang")
	}

	// Switch dengan pernyataan init, mirip if
	switch n := 4; n {
	case 0:
		fmt.Println("Nol")
	default:
		fmt.Println("Bukan nol: ", n)
	}
}

func ContohFallthrough() {
	langkah := 2
	switch langkah {
	case 1:
		fmt.Println("Persiapan bahan")
		fallthrough
	case 2:
		fmt.Println("Panaskan minyak")
		fallthrough
	case 3:
		fmt.Println("Masak")
	case 4:
		fmt.Println("Sajikan")
	}
	// Keluaran: Panaskan minyak, lalu Masak
}
