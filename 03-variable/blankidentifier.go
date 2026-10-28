package main

import (
	"fmt"
	_ "image/png" // dipanggil namun tidak digunakan
)

// function dengan dua return
func ambilData() (string, int) {
	return "Andi", 27
}

// parameter kedua sengaja dibuang (tidak dipakai)
func sapa(nama string, _ int) {
	fmt.Println("Halo", nama)
}

func BlankIndentifier() {
	// membuang hasil pertama, hanya umur yang dipakai
	_, umur := ambilData()
	fmt.Println("Umur :", umur)

	// membuang salah satu nilai dri penugasan berantai
	x, _ := 10, 20
	fmt.Println("x :", x)

	// membuang index saat iterasi
	buah := []string{"apel", "jeruk", "mangga"}
	for _, buah := range buah {
		fmt.Println(buah)
	}

	// argument yang tidak dipakai oleh fungsi
	sapa("Andi", 27)
}
