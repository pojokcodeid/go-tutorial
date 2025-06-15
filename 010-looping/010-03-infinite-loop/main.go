package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break // Breaks out of the loop when i is 5
		}
	}
}
