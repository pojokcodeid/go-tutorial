package main

import "fmt"

func main() {
	// cari angka di array dengan value 7
	angka := []int{3, 5, 7, 9, 11}

	cari := 7
	status := false

	for _, v := range angka {
		if v == cari {
			status = true
			break
		}
	}
	if !status {
		fmt.Println("Angka %v tidak ditemukan \n", cari)
	} else {
		fmt.Printf("Angka %v ditemukan \n", cari)
	}
	fmt.Println("Selesai")
}
