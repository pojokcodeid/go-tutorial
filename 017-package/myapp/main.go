// main.go
package main

import (
	"fmt"
	"mymodule/com/utils/kalkulator" // import custom package berdasarkan folder
)

func main() {
	hasil1 := kalkulator.Tambah(10, 5)
	hasil2 := kalkulator.Kurang(10, 5)

	fmt.Println("Hasil Tambah:", hasil1) // Output: 15
	fmt.Println("Hasil Kurang:", hasil2) // Output: 5
}
