package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/alifdwt/digitalent-golang/models"
)

func main() {
	students := []models.Student{
		{
			Absen:     1,
			Nama:      "Alif Dewantara",
			Alamat:    "Jl. Mangga No. 2, Bogor",
			Pekerjaan: "Full Stack Developer",
			Alasan:    "Saya ingn memperdalam skill tentang backend development",
		},
		{
			Absen:     2,
			Nama:      "Budi",
			Alamat:    "Jl. Durian No. 3, Sukabumi",
			Pekerjaan: "Programmer",
			Alasan:    "Saya ingin membuat aplikasi",
		},
		{
			Absen:     3,
			Nama:      "Caca",
			Alamat:    "Jl. Apel No. 4, Bandung",
			Pekerjaan: "Backend Engineer",
			Alasan:    "Meraih mimpi",
		},
		{
			Absen:     4,
			Nama:      "Dedi",
			Alamat:    "Jl. Anggur No. 5, Semarang",
			Pekerjaan: "Guru SMK",
			Alasan:    "Saya ingin mengajarkan bahasa pemrograman",
		},
	}

	arg := os.Args
	if len(arg) == 2 {
		absen := arg[1]
		absenInt, err := strconv.Atoi(absen)

		if err != nil {
			fmt.Println("Absen harus berupa angka")
			return
		}

		if absenInt < 1 {
			fmt.Println("Absen harus lebih dari 0")
			return
		}

		for _, student := range students {
			if student.Absen == absenInt {
				fmt.Printf("Nama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan: %s\n\n", student.Nama, student.Alamat, student.Pekerjaan, student.Alasan)
				return
			}
		}
		fmt.Println("Absen tidak ditemukan")
		return
	} else if len(arg) == 1 {
		for _, student := range students {
			if student == students[len(students)-1] {
				fmt.Printf("Absen: %d\nNama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan: %s", student.Absen, student.Nama, student.Alamat, student.Pekerjaan, student.Alasan)
			} else {
				fmt.Printf("Absen: %d\nNama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan: %s\n\n", student.Absen, student.Nama, student.Alamat, student.Pekerjaan, student.Alasan)
			}
		}
		return
	}
}
