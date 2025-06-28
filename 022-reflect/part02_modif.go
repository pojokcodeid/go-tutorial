package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4

	// Obtain a reflect.Value
	p := reflect.ValueOf(&x) // Note: pass a pointer to x
	v := p.Elem()            // Dereference the pointer to get the value

	fmt.Println("Before:", x) // Output: Before: 3.4

	// Modify the value
	if v.CanSet() {
		v.SetFloat(7.1)
	}

	fmt.Println("After:", x) // Output: After: 7.1
}
