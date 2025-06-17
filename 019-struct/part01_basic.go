package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Year   int
}

func main() {
	b := Book{
		Title:  "Belajar Go",
		Author: "Pojok Code",
		Year:   2025,
	}

	fmt.Println("Judul:", b.Title)
	fmt.Println("Penulis:", b.Author)
	fmt.Println("Tahun:", b.Year)
}
