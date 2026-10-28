package main

import (
	"fmt"
)

type Pengguna struct {
	Nama string
	Umur int
}

func KeywordNew() {
	// mengembalikan *Pengguna yang isinya zero value
	p := new(Pengguna)
	fmt.Printf("Type = %T Isi=%v \n", p, p)

	// karena sudah pointer field bisa langsung diisi
	p.Nama = "Andi"
	p.Umur = 27
	fmt.Println(*p)

	// ekuivalen tanpa new: var biasa + operator &
	p2 := &Pengguna{Nama: "Andi 2", Umur: 27}
	fmt.Println(*p2)
}
