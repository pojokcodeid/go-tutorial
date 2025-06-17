package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix)
	// Output:
	// [[1 2 3]
	//  [4 5 6]
	//  [7 8 9]]
	fmt.Println(matrix[1][1]) // Output: 5

	matrix[1][1] = 10
	fmt.Println(matrix)
}
