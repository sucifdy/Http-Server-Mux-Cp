package main

import (
	"fmt"
	"net/http"
	"time"
)

// TimeHandler untuk menampilkan waktu saat ini
func TimeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Mendapatkan waktu saat ini
		now := time.Now()
		// Mengatur format output waktu
		output := fmt.Sprintf("%s, %d %s %d", now.Weekday(), now.Day(), now.Month(), now.Year())

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(output))
	}
}

// SayHelloHandler untuk menyapa pengguna
func SayHelloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello there"))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
		}
	}
}

// GetMux untuk mengatur routing
func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	// Menambahkan route untuk /time dan /hello
	mux.HandleFunc("/time", TimeHandler())
	mux.HandleFunc("/hello", SayHelloHandler())
	return mux
}

func main() {
	// Menjalankan server pada localhost:8080
	http.ListenAndServe("localhost:8080", GetMux())
}
