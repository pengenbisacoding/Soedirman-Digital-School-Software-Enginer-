package main

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type HasilPerhitungan struct {
	Bentuk   string  `json:"bentuk"`
	JariJari float64 `json:"jari_jari,omitempty"`
	Sisi     float64 `json:"sisi,omitempty"`
	Luas     float64 `json:"luas"`
	Keliling float64 `json:"keliling"`
}

func hitungLuas(bentuk string, jariJari float64, sisi float64) float64 {
	switch bentuk {
	case "lingkaran":
		return math.Pi * math.Pow(jariJari, 2)
	case "segitiga":
		return (sisi * sisi) / 2
	case "persegi":
		return sisi * sisi
	default:
		return 0
	}
}

func hitungKeliling(bentuk string, jariJari float64, sisi float64) float64 {
	switch bentuk {
	case "lingkaran":
		return 2 * math.Pi * jariJari
	case "segitiga":
		return 3 * sisi
	case "persegi":
		return 4 * sisi
	default:
		return 0
	}
}

func hitungHandler(c *fiber.Ctx) error {
	bentuk := c.FormValue("bentuk")
	jariJari, _ := strconv.ParseFloat(c.FormValue("jari_jari"), 64)
	sisi, _ := strconv.ParseFloat(c.FormValue("sisi"), 64)

	luas := hitungLuas(bentuk, jariJari, sisi)
	keliling := hitungKeliling(bentuk, jariJari, sisi)

	hasil := HasilPerhitungan{
		Bentuk:   bentuk,
		JariJari: jariJari,
		Sisi:     sisi,
		Luas:     luas,
		Keliling: keliling,
	}

	return c.JSON(hasil)
}

func main() {
	app := fiber.New()

	app.Post("/hitung", hitungHandler)

	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
