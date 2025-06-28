package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4

	// Obtain the reflect.Type and reflect.Value
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println("Type:", t)        // Output: Type: float64
	fmt.Println("Value:", v)       // Output: Value: 3.4
	fmt.Println("Kind:", t.Kind()) // Output: Kind: float64
}
