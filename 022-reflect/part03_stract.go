package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Alice", Age: 30}

	v := reflect.ValueOf(p)

	// Check if the value is a struct
	if v.Kind() == reflect.Struct {
		fmt.Println("Struct name:", v.Type().Name())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field %d: %s %s = %v\n", i, v.Type().Field(i).Name, v.Field(i).Type(), v.Field(i).Interface())
		}
	}
}
