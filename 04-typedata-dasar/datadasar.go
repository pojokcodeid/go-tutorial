package main

import (
	"fmt"
)

func DataDasar() {
	var umur int = 27
	var tinggi float32 = 172.3
	var nama string = "Dwi"
	var status bool = true

	// %T mencetak TIPE data, %v mencetak NILAI
	fmt.Printf("%T type %v\n", umur, umur)
	fmt.Printf("%T type %v\n", tinggi, tinggi)
	fmt.Printf("%T type %v\n", nama, nama)
	fmt.Printf("%T type %v\n", status, status)

	// contoh nilai dasar
	var a int
	var b float64
	var c string
	var d bool
	var e []int
	var f map[string]int

	fmt.Printf("int :%v\n", a)
	fmt.Printf("float :%v\n", b)
	fmt.Printf("string :%v\n", c)
	fmt.Printf("bool :%v\n", d)
	fmt.Printf("slice :%v\n", e)
	fmt.Printf("map :%v\n", f)

	// byte dan rune
	s := "Go🎉"

	fmt.Println("jumlah byte :", len(s)) // 6 (emoji memakai 4 byte)

	// Iterasi per rune (karakter Unicode)
	for i, r := range s {
		fmt.Printf("posisi %d: %c (kode %d)\n", i, r, r)
	}
}
