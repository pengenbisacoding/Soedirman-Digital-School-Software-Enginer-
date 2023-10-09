package main

import "fmt"

// Membuat struct Person
type Person struct {
	Name string
	Age  int
}

// Membuat method di struct Person
func (p Person) SayHello() {
	fmt.Printf("Halo, nama saya %s dan umur saya %d tahun.\n", p.Name, p.Age)
}

// Membuat struct Car
type Car struct {
	Brand  string
	Year   int
	Engine string
}

// Membuat method di struct Car
func (c Car) Drive() {
	fmt.Printf("Saya sedang mengendarai mobil %s tahun %d dengan mesin %s.\n", c.Brand, c.Year, c.Engine)
}

func main() {
	// Membuat objek Person dan memanggil method SayHello()
	person := Person{
		Name: "Harits",
		Age:  20,
	}
	person.SayHello()

	// Membuat objek Car dan memanggil method Drive()
	car := Car{
		Brand:  "Toyota",
		Year:   2022,
		Engine: "V6",
	}
	car.Drive()
}
