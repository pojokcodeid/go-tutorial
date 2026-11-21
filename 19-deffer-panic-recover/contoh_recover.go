package main

import (
	"fmt"
	"net/http"
)

func amanHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				http.Error(w, "terjadi kesalahan internal", http.StatusInternalServerError)
				fmt.Println("panic ditangkap:", rec)
			}
		}()
		fn(w, r)
	}
}

func PanggilRoot() {
	http.HandleFunc("/", amanHandler(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Selamat datang di halaman utama!")
	}))

	http.ListenAndServe(":8080", nil)
}
