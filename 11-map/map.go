package main

import (
	"fmt"
)

func MembuatMap() {
	// membuat map kosong
	umur := make(map[string]int)

	// menambah atau mengubah entri
	umur["Andi"] = 27
	umur["Budi"] = 30
	umur["Andi"] = 25 // akan menimpa nilai lama

	// Map dengan Literal
	stok := map[string]int{
		"beras": 50,
		"gula":  20,
	}

	// menghapus entri
	delete(stok, "gula")
	fmt.Println(umur)
	fmt.Println(stok["beras"])
	fmt.Println(len(stok))
}

func PengecekanKey() {
	nilai := map[string]int{"matematika": 90}
	// key tidak ditemukan ok: tidak bisa membedakan key kosong dengan nilai 0
	// karena samasama 0
	fmt.Println(nilai["fisika"])

	// dengan key ada bernilai boolean
	skor, ada := nilai["fisika"]
	if ada {
		fmt.Println("Fisika:", skor)
	} else {
		fmt.Println("Fisika belum ada")
	}

	// pengecekan dalam satu baris
	if skor, ada := nilai["matematika"]; ada {
		fmt.Println("Matematika:", skor)
	}
}

func Iterasi() {
	// Slice: i = indeks, v=nilai
	warna := []string{"merah", "hijau", "biru"}
	for i, w := range warna {
		fmt.Println(i, w)
	}

	// Map: k=key, v=value (urutan TIDAK dijamin)
	umur := map[string]int{"Andi": 27, "Budi": 30}
	for nama, u := range umur {
		fmt.Println(nama, "berumur", u)
	}

	// Abaikan salah satu dengan _
	for _, u := range umur {
		fmt.Println("umur:", u)
	}
}

func MapSlice() {
	kota := map[string][]string{}
	kota["Bandung"] = append(kota["Bandung"], "Andi")
	kota["Bandung"] = append(kota["Bandung"], "Citra")
	kota["Surabaya"] = append(kota["Surabaya"], "Budi")

	for nama, orang := range kota {
		fmt.Printf("%s (%d orang): %v", nama, len(orang), orang)
	}
}
