package main

import (
	"fmt"
	"sync"
	"time"
)

func tugas(nama string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(nama, "-> langkah", i)
		time.Sleep(100 * time.Microsecond)
	}
}

func Panggil1() {
	// jalankan sebagai go routine dengan kata kunci "go"
	go tugas("A")
	go tugas("B")

	// Beri waktu agar goroutine menyelesaikan pekerjaannya.
	// Nanti kita ganti dengan sync.WaitGroup yang lebih rapi.
	time.Sleep(time.Second)
	fmt.Println("selesai")
}

// contoh masalah race condition
func RaceCondition() {
	counter := 0
	var wg sync.WaitGroup

	for i := 0; i <= 1000; i++ {
		wg.Go(func() {
			counter++ // race condition: goroutine menulis bersamaan
		})
	}
	wg.Wait()
	fmt.Println("counter =", counter) // hasilnya tidak pasti, sering < 1000
}
