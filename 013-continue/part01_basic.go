package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println("Angka:", i)
	}
}
