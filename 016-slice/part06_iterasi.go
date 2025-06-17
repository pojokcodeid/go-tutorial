package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
