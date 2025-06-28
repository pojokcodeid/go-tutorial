package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Eve", Age: 40}

	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData)) // Output: {"Name":"Eve","Age":40}
}
