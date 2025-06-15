package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				goto end
			}
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}
end:
	fmt.Println("Exited the nested loop")
}
