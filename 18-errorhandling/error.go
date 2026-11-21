package main

import (
	"errors"
	"fmt"
)

func bagi(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Pembagi tidak boleh nol")
	}
	return a / b, nil
}

func ContohError() {
	hasil, err := bagi(10, 0)
	if err != nil {
		fmt.Println("Terjadi kesalahan:", err)
		return
	}
	fmt.Println("Hasil pembagian:", hasil)

	// fmt.Errorf untuk pesan berformat
	nama := "budi"
	err2 := fmt.Errorf("Nama %s tidak ditemukan", nama)
	fmt.Println(err2)
}

// contoh membungkus error dengan %w
var ErrTidakDitemukan = errors.New("data tidak ditemukan")

func cari(id int) (string, error) {
	if id != 1 {
		// %w membungkus error asli agar bisa dilacak
		return "", fmt.Errorf("cari data id %d: %w", id, ErrTidakDitemukan)
	}
	return "Andi", nil
}

func ContohError2() {
	_, err := cari(99)
	if err != nil {
		fmt.Println("Terjadi kesalahan:", err)
		// error.Is memeriksan apakah rantai error memuat target
		if errors.Is(err, ErrTidakDitemukan) {
			fmt.Println("-> memang error'tidak ditemukan'")
		}
	}
}

// contoh custom error
type ValidationError struct {
	Field string
	Pesan string
}

// implementasi interface error
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validasi gagal pada %q: %s", e.Field, e.Pesan)
}

func validasiUmur(umur int) error {
	if umur < 0 {
		return &ValidationError{Field: "umur", Pesan: "tidak boleh negatif"}
	}
	if umur > 150 {
		return &ValidationError{Field: "umur", Pesan: "tidak realistis"}
	}
	return nil
}

func ContohError3() {
	err := validasiUmur(-5)
	if err != nil {
		fmt.Println(err)

		// errors.AsType mengambil error bertipe tertentu secara langsung dan type-safe
		if ve, ok := errors.AsType[*ValidationError](err); ok {
			fmt.Println("-> field brmasalah:", ve.Field)
		}
	}
}
