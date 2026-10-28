package main

import (
	"errors"
	"fmt"
	"strconv"
)

func Kondisional() {
	nilai := 78
	if nilai >= 85 {
		fmt.Println("Nilai A")
	} else if nilai >= 70 {
		fmt.Println("Nilai B")
	} else if nilai >= 60 {
		fmt.Println("Nilai C")
	} else {
		fmt.Println("Nilai D")
	}
}

func Initialisasi() {
	if angka, err := strconv.Atoi("123"); err == nil {
		fmt.Println("Berhasil di konversi:", angka)
	} else {
		fmt.Println("gagal konversi:", err)
	}
	// Sama tanpa error, memakai pembatas lebar
	if n := 15; n%2 == 0 {
		fmt.Println(n, "genap")
	} else {
		fmt.Println(n, "ganjil")
	}
}

func ambilStok(produk string) (int, error) {
	if produk == "" {
		return 0, errors.New("Nama product kosong")
	}
	return 42, nil
}

func CekError() {
	stok, err := ambilStok("")
	if err != nil {
		fmt.Println("Terjadi error:", err)
		return
	}
	fmt.Println("Stok:", stok)
}
