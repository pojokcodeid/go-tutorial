package main

import (
	"errors"
	"fmt"
)

// satu nilai kembalian
func tambah(a int, b int) int {
	return a + b
}

// parameter bertype sama boleh disingkat
func kali(a, b int) int {
	return a * b
}

// fungsi tanpa kembalian sama sekali
func sapa(a string) {
	fmt.Println("Halo", a)
}

func FungsiDasar() {
	fmt.Println(tambah(1, 2))
	fmt.Println(kali(1, 2))
	sapa("Budi")
}

// fungsi mengembalikan multiple
func bagi(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Tidak bisa membagi dengan nol")
	}
	return a / b, nil
}

// mengembalikan 3 nilai min, max, rata-rata
func statustic(data []int) (int, int, float64) {
	total := 0
	min, max := data[0], data[0]
	for _, v := range data {
		total += v
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max, float64(total) / float64(len(data))
}

func MultiReturn() {
	hasil, err := bagi(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Hasil:", hasil)
	}

	lo, hi, avg := statustic([]int{5, 2, 9, 1})
	fmt.Println(lo, hi, avg)
}

// Nama hasil & err sudah "terdeklarasi" di tanda tangan.
func bagi2(a, b float64) (hasil float64, err error) {
	if b == 0 {
		err = fmt.Errorf("pembagi nol")
		return // bare return -> mengembalikan hasil & err
	}
	hasil = a / b
	return // bare return
}

// Kadang lebih jelas bila memakai return eksplisit
func panjangDanLebar() (p int, l int) {
	p = 10
	l = 5
	return p, l
}

func NamedReturn() {
	h, err := bagi2(8, 2)
	fmt.Println(h, err)
	fmt.Println(panjangDanLebar())
}
