package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{5, 9, 3, 1, 7} // Data bilangan yang akan diurutkan

	fmt.Println("Sebelum diurutkan:", numbers)

	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))

	fmt.Println("Setelah diurutkan:", numbers)

	// Membuat slice dengan jumlah data 5
	slice := make([]int, 5)

	// Menambahkan data ke dalam slice sebanyak 3
	for i := 0; i < 3; i++ {
		var data int
		fmt.Printf("Masukkan data ke-%d: ", i+1)
		fmt.Scan(&data)
		slice = append(slice, data)
	}

	// Menampilkan isi slice
	fmt.Println("Isi slice:")
	fmt.Println(slice)
}
