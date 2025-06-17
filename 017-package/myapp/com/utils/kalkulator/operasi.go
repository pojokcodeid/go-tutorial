// kalkulator/operasi.go
package kalkulator

// Fungsi Tambah diekspor (bisa diakses dari luar package)
func Tambah(a, b int) int {
	return a + b
}

// Fungsi Kurang juga diekspor
func Kurang(a, b int) int {
	return a - b
}
