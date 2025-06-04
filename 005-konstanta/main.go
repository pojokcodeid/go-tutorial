package main

import (
	"fmt"
)

const PI = 3.14

// contoh multiple constanta
const (
	A int = 1
	B     = 3.14
	C     = "Hi!"
)

func main() {
	// tidak bisa di ubah nilai PI
	// PI = 10
	fmt.Println(PI)
	// cetak multiple constanta
	// prilakunya sama dengan single constata yaotu tidak dapat diubah
	fmt.Println(A, B, C)
}
