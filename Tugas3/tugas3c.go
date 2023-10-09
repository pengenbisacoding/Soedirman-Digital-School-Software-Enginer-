package main

import "fmt"

// Membuat struct Student
type Student struct {
	Name  string
	Grade int
}

// Membuat function untuk mengubah data Grade menggunakan pointer
func ChangeGrade(s *Student, newGrade int) {
	s.Grade = newGrade
}

func main() {
	// Membuat objek Student
	student := Student{
		Name:  "John",
		Grade: 80,
	}

	// Menampilkan data awal
	fmt.Printf("Nama: %s, Grade: %d\n", student.Name, student.Grade)

	// Mengubah data Grade menggunakan pointer
	ChangeGrade(&student, 90)

	// Menampilkan data setelah diubah
	fmt.Printf("Nama: %s, Grade: %d\n", student.Name, student.Grade)
}
