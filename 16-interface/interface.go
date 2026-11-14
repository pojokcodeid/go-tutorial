package main

import (
	"fmt"
	"math"
)

// interface: kumpulan method yang harus dimiliki
type BangunDatar interface {
	Luas() float64
	Keliling() float64
}

type Persegi struct {
	Sisi float64
}
type Lingkaran struct{ R float64 }

// peersegi memiliki luas dan keliling otomatis memenuhi bangun datar
func (p Persegi) Luas() float64 {
	return p.Sisi * p.Sisi
}

func (p Persegi) Keliling() float64 {
	return 4 * p.Sisi
}

// lingkaran juga memenuhi bangun datar
func (l Lingkaran) Luas() float64 {
	return math.Pi * l.R * l.R
}

func (l Lingkaran) Keliling() float64 {
	return 2 * math.Pi * l.R
}

func cetak(b BangunDatar) {
	fmt.Printf("Luas=%.2f Keliling=%.2f \n", b.Luas(), b.Keliling())
}

func PanggilCetak() {
	cetak(Persegi{Sisi: 4})
	cetak(Lingkaran{R: 7})
}

/* contoh interface implicit */

type Penyipanan interface {
	Simpan(kunci, nilai string)
	Ambil(kunci string) string
}

// implementasi memori
type Memori map[string]string

func (m Memori) Simpan(k, v string)    { m[k] = v }
func (m Memori) Ambil(k string) string { return m[k] }

// Implementasi "Palsu" yang hanya mencatat ke layar
type LogPenyimpanan struct{}

func (LogPenyimpanan) Simpan(k, v string) {
	fmt.Println("[LOG] menumpan", k, "=", v)
}

func (LogPenyimpanan) Ambil(k string) string {
	fmt.Println("[LOG] mengambil", k)
	return ""
}

// fungsi ini tidak peduli implementasi mana yang dipakai
func pakai(s Penyipanan) {
	s.Simpan("nama", "Andi")
	fmt.Println("hasil:", s.Ambil("nama"))
}

func PanggilDependencyInversion() {
	pakai(Memori{})
	pakai(LogPenyimpanan{})
}

// contoh empty interface
// menerima nilai tipe apapun
// any adalah alias dari interface{}
// fungsi yang menerima any dapat di beri nilai apapun
func tampilkan(x any) {
	fmt.Printf("nilai=%v type=%T", x, x)
}

func EmptyInterface() {
	tampilkan(42)
	tampilkan("hallo")
	tampilkan(true)
	tampilkan(map[string]int{"a": 1})
	tampilkan([]int{1, 2, 3})

	// slice campuran
	data := []any{1, "dua", 3.0, false}
	for _, d := range data {
		fmt.Println(d)
	}
}
