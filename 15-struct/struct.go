package main

import (
	"fmt"
)

type Pengguna struct {
	Nama  string
	Umur  int
	Aktif bool
}

func BasicStruct() {
	// degnan nama field paling jelas
	u1 := Pengguna{Nama: "Andi", Umur: 27, Aktif: true}

	// berurutan sesuai nama field
	u2 := Pengguna{"Budi", 31, false}

	// deklarasi lalu isi field
	var u3 Pengguna
	u3.Nama = "Citra"
	u3.Umur = 24

	fmt.Println(u1)
	fmt.Println(u2.Nama, u2.Umur)
	fmt.Println(u3)

	// field yang kosong otomatis zero value
	fmt.Println(u3.Aktif) // zero value bool adalah false
}

// ini contoh struct nested dan anomimouse
type Alamat struct {
	Jalan string
	Kota  string
}

// nested struct (struk didalam struck)
type Karyawan struct {
	Nama   string
	Alamat Alamat
}

func NestedAnonim() {
	k := Karyawan{
		Nama: "Dewi",
		Alamat: Alamat{
			Jalan: "Jl Jurago no 20",
			Kota:  "jakarta",
		},
	}

	fmt.Println(k.Nama, k.Alamat.Kota)

	// anonimouse struk : struck tanpa nama type cocok untuk data sekali pakai
	titik := struct {
		X, Y int
	}{Y: 3, X: 5}
	fmt.Println(titik)
}

// contoh embeding struck
type Hewan struct {
	Nama string
}

func (h Hewan) Suara() string {
	return h.Nama + " bersuara...."
}

type Kucing struct {
	Hewan // embeding struck
	Jenis string
}

func EmbedingStruck() {
	k := Kucing{
		Hewan{Nama: "Milo"},
		"Persia",
	}
	fmt.Println(k.Suara())
	fmt.Println(k.Nama)
}

// value receiver
type Counter struct {
	Nilai int
}

// Value receiver -> menerima SALINAN (tidak mengubah asli)
func (c Counter) Tampilkan() {
	fmt.Println("nilai:", c.Nilai)
}

// Pointer receiver -> bisa mengubah struck asli
func (c *Counter) Tambah(n int) {
	c.Nilai += n
}

func ContohRecivier() {
	c := Counter{Nilai: 10}
	c.Tampilkan() // 10
	c.Tambah(5)   // go otomatis mengirim ke &c
	c.Tampilkan() //15

	p := &Counter{Nilai: 100}
	p.Tambah(1)   // juga bisua dari pointer
	p.Tampilkan() // 101
}
