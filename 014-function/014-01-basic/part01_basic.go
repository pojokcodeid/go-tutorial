package main

import "fmt"

// function tanpa parameter
func sayHello() {
	fmt.Println("Halo World")
}

// Ini function dengan parameter
func sapa(nama string) {
	fmt.Println("Halo", nama)
}

// function dengan return
func getGreeting(nama string) string {
	return "Halo " + nama
}

// function degnan return multiple
func swap(a, b int) (int, int) {
	return b, a
}

// function dengan nama return
func divide(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

func main() {
	sayHello()
	sapa("Asep")
	sapa("Budi")
	fmt.Println(getGreeting("Caca"))

	x, y := swap(1, 2)
	fmt.Println(x, y)

	quotient, remainder := divide(5, 2)
	fmt.Println(quotient, remainder)
}
