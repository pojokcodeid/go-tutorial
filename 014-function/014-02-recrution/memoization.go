package main

import "fmt"

// Buat map untuk menyimpan hasil yang sudah dihitung
var memo = map[int]int{}

func fibonacci(n int) int {
	// Cek apakah sudah pernah dihitung
	if val, found := memo[n]; found {
		return val
	}

	// Base case
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	// Hitung dan simpan ke memo
	result := fibonacci(n-1) + fibonacci(n-2)
	memo[n] = result
	return result
}

func main() {
	for i := 0; i < 20; i++ {
		fmt.Print(fibonacci(i), " ")
	}
}
