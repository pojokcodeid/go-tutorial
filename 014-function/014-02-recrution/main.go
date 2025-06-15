package main

import "fmt"

func hitungMundur(n int) {
	if n == 0 {
		fmt.Println("Selesai!")
		return
	}

	fmt.Println("Hitung:", n)
	hitungMundur(n - 1) // function memanggil dirinya sendiri
}

func factorial(n int) int {
	if n == 0 {
		return 1 // Base case
	}
	return n * factorial(n-1) // Recursive case
}

// fibonacci
// fibonaci adalah jumlah dari dua angka sebelumnya
// 0, 1, 1, 2, 3, 5, 8, 13, 21, ...
func fib(n int) int {
	if n == 0 {
		return 0 // Base case
	} else if n == 1 {
		return 1 // Base case
	}
	return fib(n-1) + fib(n-2) // Recursive case
}

func main() {
	hitungMundur(3)
	fmt.Println(factorial(5))
	fmt.Println(fib(10))
}
