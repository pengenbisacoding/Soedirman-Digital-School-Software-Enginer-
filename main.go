package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Program Menghitung Luas Persegi, Lingkara, Segitiga")

	// Menghitung luas persegi
	sisiPersegi := 5.0                      // Menginisialisasi variabel sisiPersegi dengan nilai 5.0
	luasPersegi := math.Pow(sisiPersegi, 2) // Menghitung luas persegi dengan menggunakan fungsi math.Pow untuk mengkuadratkan sisiPersegi
	fmt.Printf("Luas Persegi dengan sisi %.2f adalah %.2f\n", sisiPersegi, luasPersegi)

	// Menghitung luas lingkaran
	jariJari := 4.0                                  // Menginisialisasi variabel jariJari dengan nilai 4.0
	luasLingkaran := math.Pi * math.Pow(jariJari, 2) // Menghitung luas lingkaran dengan menggunakan konstanta math.Pi dan fungsi math.Pow untuk mengkuadratkan jariJari
	fmt.Printf("Luas Lingkaran dengan jari-jari %.2f adalah %.2f\n", jariJari, luasLingkaran)

	// Menghitung luas segitiga
	alas := 6.0                         // Menginisialisasi variabel alas dengan nilai 6.0
	tinggi := 8.0                       // Menginisialisasi variabel tinggi dengan nilai 8.0
	luasSegitiga := 0.5 * alas * tinggi // Menghitung luas segitiga dengan perkalian 0.5, alas, dan tinggi
	fmt.Printf("Luas Segitiga dengan alas %.2f dan tinggi %.2f adalah %.2f\n", alas, tinggi, luasSegitiga)
	// D. Menghitung Energi Potensial dan Kinetik
	fmt.Println("Program Menghitung Energi Potensial dan Kinetik")
	// Menghitung energi potensial
	massa := 5.0       // Massa benda dalam kg
	gravitasi := 9.8   // Percepatan gravitasi dalam m/s^2
	ketinggian := 10.0 // Ketinggian benda dalam meter

	energiPotensial := massa * gravitasi * ketinggian
	fmt.Printf("Energi Potensial: %.2f Joule\n", energiPotensial)

	// Menghitung energi kinetik
	kecepatan := 15.0 // Kecepatan benda dalam m/s

	energiKinetik := 0.5 * massa * kecepatan * kecepatan
	fmt.Printf("Energi Kinetik: %.2f Joule\n", energiKinetik)

	// E. Menghitung frekuensi atau getaran

	fmt.Println("Program Menghitung Frekuensi atau Getaran")

	// Input periode dalam detik
	var periode float64
	fmt.Print("Masukkan periode dalam detik: ")
	fmt.Scanln(&periode)

	// Menghitung frekuensi
	frekuensi := 1 / periode
	fmt.Printf("Frekuensi: %.2f Hz\n", frekuensi)

	// Menghitung jumlah getaran dalam satu menit
	jumlahGetaran := 60 / periode
	fmt.Printf("Jumlah Getaran dalam Satu Menit: %.2f\n", jumlahGetaran)

	fmt.Println("Program Konversi Suhu")
	// F. Konversi untuk semua satuan suhu
	// Input suhu dalam Celcius
	var suhuCelcius float64
	fmt.Print("Masukkan suhu dalam Celcius: ")
	fmt.Scanln(&suhuCelcius)

	// Konversi ke Fahrenheit
	suhuFahrenheit := (suhuCelcius * 9 / 5) + 32
	fmt.Printf("Suhu dalam Fahrenheit: %.2f\n", suhuFahrenheit)

	// Konversi ke Kelvin
	suhuKelvin := suhuCelcius + 273.15
	fmt.Printf("Suhu dalam Kelvin: %.2f\n", suhuKelvin)

	// Konversi ke Reamur
	suhuReamur := suhuCelcius * 4 / 5
	fmt.Printf("Suhu dalam Reamur: %.2f\n", suhuReamur)

	// Konversi ke Rankine
	suhuRankine := (suhuCelcius + 273.15) * 9 / 5
	fmt.Printf("Suhu dalam Rankine: %.2f\n", suhuRankine)
}
