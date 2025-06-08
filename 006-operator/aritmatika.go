package main

import (
	"fmt"
)

var value = (((2+6)%3)*4 - 2) / 3

func main() {
	var (
		sum1 = 100 + 50    // 150 (100 + 50)
		sum2 = sum1 - 250  // 400 (150 + 250)
		sum3 = sum2 + sum2 // 800 (400 + 400)
	)
	fmt.Println(sum3)
	fmt.Println(value)
}
