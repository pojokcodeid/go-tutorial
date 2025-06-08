package main

import "fmt"

func main() {
	// type data non desimal
	var bilanganPositif uint8 = 80
	var bilanganNegatif = -1243423644 // compiler secara cerdas otomatis menentukan int32

	fmt.Println("Bilangan Positif :",bilanganPositif)
	fmt.Println("Bilangan Negatif :",bilanganNegatif)
}
