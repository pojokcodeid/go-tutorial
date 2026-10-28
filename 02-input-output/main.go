package main

import (
	"fmt"
)

type Product struct {
	Nama  string
	Harga int
}

func main() {
	//  print tidak menambahkan new line sendiri
	fmt.Print("Helo")
	fmt.Println(" World") // menabhakan new line diakhir
	//  Println akan memberikan spasi antar argument + new line diakhir
	fmt.Println("Umur:", 27, "active:", true)

	//  perbedaan aturan spasi pada print
	fmt.Print("a", "b") // ab dua string -> tanpa spasi
	fmt.Println()
	fmt.Print(1, "a") // string tanpa spasi
	fmt.Println()
	fmt.Print(1, 2) // 1 2 dua angka di beri spasi
	fmt.Println()

	p := Product{Nama: "Kopi", Harga: 2000}
	fmt.Printf("int :%d \n", p.Harga)
	fmt.Printf("string : %s kutipan: %q \n", p.Nama, p.Nama)
	fmt.Printf("boolean:%t | float: %.2f | hex: %x \n", true, 3.1423456, 255)
	fmt.Printf("default: %v \n", p)
	fmt.Printf("rinci: %+v \n", p)
	fmt.Printf("Type: %T | persentase: 100%% \n", p)
	fmt.Printf("lebar :[%5d] [%-5d]", 42, 42)
	fmt.Println()

	// Sprintf mengembalikan string tanpa mencetak
	nama := "Andi"
	umur := 27
	s := fmt.Sprintf("Nama :%s Umur:%d", nama, umur)
	fmt.Println(s)
	fmt.Println("Panjang :", len(s))
}
