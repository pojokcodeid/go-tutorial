package main

import "fmt"

func main() {
	// type data string
	var name string = "Golang"
	fmt.Println("Name : ", name)

	// string dengan backtick
	var name2 string = `Golang \n "Golang" dan 'Golang'`
	fmt.Println("Name 2 : ", name2)

	// string multiline
	var name3 string = `Golang
	Golang
	Golang`
	fmt.Println("Name 3 : ", name3)

	// petik dua dengan baris baru tidak dijinkan
	var name4 string = "Golang \n Golang \n Golang"
	fmt.Println("Name 4 : ", name4)
}
