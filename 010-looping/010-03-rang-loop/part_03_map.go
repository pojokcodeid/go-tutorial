package main

import "fmt"

func main() {
	myMap := map[string]int{"one": 1, "two": 2, "three": 3}
	for key, value := range myMap {
		fmt.Println("Key:", key, "Value:", value)
	}
}
