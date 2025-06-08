package main

import "fmt"

func main() {
	// contoh input scanner
	var name string
	var age int
	fmt.Print("Masukkan nama dan umur: ")
	fmt.Scan(&name, &age)
	fmt.Println("Nama:", name, "Umur:", age)

	// contoh input scanner
	var name string
	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&name)
	fmt.Println("Nama:", name)

	// contoh input scanf
	var name string
	var age int
	fmt.Print("Masukkan nama dan umur (format: Nama Umur): ")
	fmt.Scanf("%s %d", &name, &age)
	fmt.Println("Nama:", name, "Umur:", age)

}
