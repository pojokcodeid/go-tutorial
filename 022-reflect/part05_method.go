package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
	p := Person{Name: "Diana", Age: 28}

	v := reflect.ValueOf(p)

	method := v.MethodByName("Greet")
	if method.IsValid() {
		method.Call(nil) // Call the method with no arguments
	}
}
