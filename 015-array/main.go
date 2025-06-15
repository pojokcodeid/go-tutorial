package main

import "fmt"

func main() {
	var numbers [5]int
	fmt.Println(numbers) // Output: [0 0 0 0 0]

	// array dengan initialize
	numbers2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers2) // Output: [1 2 3 4 5]

	// panjang array dinamis
	numbers3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(numbers3) // Output: [1 2 3 4 5]

	// mengacess element array
	fmt.Println(numbers3[0]) // Output: 1
	fmt.Println(numbers3[4]) // Output: 5

	// mengisi array
	numbers3[2] = 10
	fmt.Println(numbers3) // Output: [1 2 10 4 5]

	// panjang array
	fmt.Println(len(numbers3)) // Output: 5

	// interasi array degnan for loop
	for i := 0; i < len(numbers3); i++ {
		fmt.Println(numbers3[i])
	}

	// For-Range loop
	for index, value := range numbers3 {
		fmt.Println(index, value)
	}

	// array multidimensi
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix)
	// access array multi dimensi
	fmt.Println(matrix[1][1]) // Output: 5

	matrix[1][1] = 10
	fmt.Println(matrix)
}
