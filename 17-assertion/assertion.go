package main

import (
	"encoding/json"
	"fmt"
)

func ContohAssertion() {
	var x any = "Hello"

	// cara aman coma-ok
	s, ok := x.(string)
	if ok {
		fmt.Println("Ini string:", s, "panjang:", len(s))
	}

	// gagal ok bersisi false, s zero value
	n, ok := x.(int)
	fmt.Println(n, ok)
}

// menjumlahkan element aapun yang bisa dianggap angka
func total(data []any) float64 {
	var hasil float64
	for _, v := range data {
		switch n := v.(type) {
		case int:
			hasil += float64(n)
		case float64:
			hasil += n
		case string:
			fmt.Println("Melewati string:", n)
		default:
			fmt.Printf("Type tidakdidukung: %T\n", n)
		}
	}
	return hasil
}

// cetak total dari data
func CetakTotal() {
	data := []any{10, 2.5, "tiga", true}
	fmt.Println("Total:", total(data))
}

func JsonDinamis() {
	text := `{"nama":"Andi","umur":27,"hobi":["membaca","berenang"]}`
	var data map[string]any
	if err := json.Unmarshal([]byte(text), &data); err != nil {
		fmt.Println("Gagal decode json:", err)
		return
	}

	for kunci, nilai := range data {
		switch v := nilai.(type) {
		case string:
			fmt.Println(kunci, "adalah string:", v)
		case float64:
			fmt.Println(kunci, "adalah float64:", v)
		case []any:
			fmt.Println(kunci, "adalah slice:", v)
		default:
			fmt.Println(kunci, "memiliki type tidak diketahui")
		}
	}
}
