package main

import (
	"fmt"
)

func ContohPerulangan() {
	for i := 0; i <= 5; i++ {
		fmt.Println("Nilai:", i)
	}
	fmt.Println("+-----------------------------------------")

	// penjumlahan 1 - 10
	var total int
	for i := 1; i <= 10; i++ {
		total += i
	}
	fmt.Println("Total:", total)
	fmt.Println("+-----------------------------------------")

	// hanya kondisi seperti while
	n := 5
	for n > 0 {
		fmt.Println(n)
		n--
	}
	fmt.Println("+-----------------------------------------")

	// loop tak terbatas + break
	hitung := 0
	for {
		fmt.Println("Hitung:", hitung)
		hitung++
		if hitung == 10 {
			break
		}
	}
	fmt.Println("+-----------------------------------------")

	// continue melewati iterasi saat ini
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // lewati bilangan genap
		}
		fmt.Println("Nilai:", i)
	}
	fmt.Println("+-----------------------------------------")

	// label untuk keluar dari loop bersarang
luar:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println("i:", i, "j:", j)
			if i == 2 {
				break luar
			}
		}
	}
	fmt.Println("+-----------------------------------------")

	// for range
	angka := []int{10, 20, 30, 40, 50}
	for i, value := range angka {
		fmt.Println(i, value)
	}
	// mengabaikan index dengan _
	for _, value := range angka {
		fmt.Println(value * 2)
	}
	fmt.Println("+-----------------------------------------")
}
