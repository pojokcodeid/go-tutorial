package main

import "fmt"

// Defining nested structs
type Address struct {
	City    string
	ZipCode int
}

type Person3 struct {
	Name    string
	Age     int
	Address Address
}

func main() {
	// Initializing a struct with nested structs
	p := Person3{
		Name: "Frank",
		Age:  45,
		Address: Address{
			City:    "New York",
			ZipCode: 10001,
		},
	}
	fmt.Println(p)
	// Output: {Frank 45 {New York 10001}}
}
