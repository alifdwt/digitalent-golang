# Digitalent Kominfo Golang 004 - Assignment 01

> This is students list manager, built using the [Go programming language](https://golang.org). It allows user to get all students data, and to get individual student data by 'absen'.

## 🧰 Installation

1. Clone the repository

```bash
git clone https://github.com/alifdwt/digitalent-golang.git
```

2. Run the application

```bash
go run main.go
```

## 📝 Documentation

### Print all students

```bash
go run main.go
```

Result:

```bash
Absen: 1
Nama: Alif Dewantara
Alamat: Jl. Mangga No. 2, Bogor
Pekerjaan: Full Stack Developer
Alasan: Saya ingn memperdalam skill tentang backend development

Absen: 2
Nama: Budi
Alamat: Jl. Durian No. 3, Sukabumi
Pekerjaan: Programmer
Alasan: Saya ingin membuat aplikasi

Absen: 3
Nama: Caca
Alamat: Jl. Apel No. 4, Bandung
Pekerjaan: Backend Engineer
Alasan: Meraih mimpi

Absen: 4
Nama: Dedi
Alamat: Jl. Anggur No. 5, Semarang
Pekerjaan: Guru SMK
Alasan: Saya ingin mengajarkan bahasa pemrograman%
```

### Print student by absen

```bash
go run main.go {absen}
```

Example:

```bash
go run main.go 2
```

Result:

```bash
Absen: 2
Nama: Budi
Kelas: Jl. Durian No. 3, Sukabumi
Pekerjaan: Programmer
Alasan: Saya ingin membuat aplikasi
```

This is an assignment for "Digitalent - Scalable Web Service with Golang" course.
