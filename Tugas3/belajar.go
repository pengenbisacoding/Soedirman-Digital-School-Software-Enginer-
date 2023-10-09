package main

import "fmt"

func main() {
	// Membuat MAP dengan tipe data string sebagai kunci (key) dan tipe data int sebagai nilai (value)
	data := map[string]int{
		"apel":   10,
		"jeruk":  5,
		"mangga": 8,
	}

	// Menampilkan data pada MAP
	for k, v := range data {
		fmt.Printf("Kunci: %s, Nilai: %d\n", k, v)
	}
}
