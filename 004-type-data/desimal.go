package main

import "fmt"

func main() {
	// type data desimal
	var bilanganPositif float32 = 80.5
	var bilanganNegatif = -1243423644.5 // compiler secara cerdas otomatis menentukan float64

	fmt.Println("Bilangan Positif :",bilanganPositif)
	fmt.Println("Bilangan Negatif :",bilanganNegatif)
}