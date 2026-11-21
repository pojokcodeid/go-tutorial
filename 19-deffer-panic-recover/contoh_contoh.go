package main

import (
	"fmt"
	"os"
)

func BacaBerkas(nama string) {
	// buka berkas lalu langsung jadualkan penutupannya
	f, err := os.Open(nama)
	if err != nil {
		fmt.Println("gagal membuka :", err)
		return
	}
	defer f.Close() // dialankan otomatis saat fungsi selesai
	// process membca berkas
	fmt.Println("berkas dibuka :", f.Name())
}

// contoh defer yang di eksekusi secara lifo
func DeferLIFO() {
	// defer dijalankan terbalik (last in, first out)
	defer fmt.Println("Ketiga")
	defer fmt.Println("Kedua")
	defer fmt.Println("Pertama")
	fmt.Println("badan fungsi")

	// nilai argument defer dibekukan saat defer ditulis
	n := 10
	defer fmt.Println("nilai n :", n)
	n = 99
	fmt.Println("nilai n diubah :", n)
}

// recover hanya bekerja di dalam defer
func amanJalnakan() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("pulih dari panic", r)
		}
	}()

	fmt.Println("mulai")
	panic("Ada yang salah fatal")
}

func PanggilPanic() {
	amanJalnakan()
	fmt.Println("Program tetap berjalan")
}
