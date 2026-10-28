package main

import (
	"fmt"
)

func main() {
	var nama string
	var umur int

	n, err := fmt.Scan(&nama, &umur)
	fmt.Printf("Scan : %d nilai, error=%v nama=%s umur=%d \n", n, err, nama, umur)

	var kota string
	var skor float64
	n2, err2 := fmt.Scanf("%s %f", &kota, &skor)
	fmt.Println("Scnf : %d nilai, error=%v %s/%.1f\n", n2, err2, kota, skor)

	var a, b int
	n3, err3 := fmt.Scanln(&a, &b)
	fmt.Println("Scanln : %d error:%v %d/%d\n", n3, err3, a, b)
}
