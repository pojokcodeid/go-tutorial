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
	p := Person{Name: "Bob", Age: 25}

	v := reflect.ValueOf(&p).Elem() // Note: pass a pointer to p

	if v.Kind() == reflect.Struct {
		nameField := v.FieldByName("Name")
		if nameField.CanSet() && nameField.Kind() == reflect.String {
			nameField.SetString("Charlie")
		}

		ageField := v.FieldByName("Age")
		if ageField.CanSet() && ageField.Kind() == reflect.Int {
			ageField.SetInt(35)
		}
	}

	fmt.Println(p) // Output: {Charlie 35}
}
