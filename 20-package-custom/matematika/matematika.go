// berkas : matematika/matematika.go
package matematika

// fugnsi tambah di export di deklarasikan dengan huruf kapital
func Tambah(a, b int) int {
	return a + b
}

// kali di tulis dengan huruf kecuil, karena funcgsi hanya diacess disini saja
func kali(a, b int) int {
	return a * b
}

func KaliDua(a int) int {
	return kali(a, 2)
}
