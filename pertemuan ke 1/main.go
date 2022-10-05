package main

import "fmt"

func main() {

	//print
	//kalau menggunnakan println auto menambahkan spasi pada setiap println
	fmt.Println("Hello world!")
	fmt.Println("Hello world!", "tes", "tes1")

	//kalau hanya print dia akan menampilkan print selanjutnya di samping
	fmt.Print("Hello world!")
	fmt.Print("Hello world!", "tes", "tes1 \n")
	//end print

	//variable

	//variable with data type 1
	var hello string = "Hello world!"
	var age int = 21

	fmt.Println(hello, age)

	//variable with data type 2
	var nama string
	var umur int

	nama = "Alifian"
	umur = 21

	fmt.Println(nama)
	fmt.Println(umur)

	//variable without data type
	var name = "Anto"
	var umurs = 20

	fmt.Printf("%T, %T", name, umurs, "\n") //printf digunakan untuk mmengecek type data pada variable dengan di tambahkan verb %T

	//variable without data type short declaration
	hp := "xiaomi"
	ram := 6

	fmt.Printf("%T, %T", hp, ram, "\n")

	//multiple variable declarations
	var nama1, nama2, nama3 string = "jono", "jana", "jini"
	var angka1, angka2, angka3 int

	angka1, angka2, angka3 = 1, 2, 3

	fmt.Println(nama1, nama2, nama3)
	fmt.Println(angka1, angka2, angka3)

	//underscore variable (untuk mengatasi error jika variable tidak digunakan)
	var tono string = "pelari"
	var tini string = "penyanyi"

	_, _ = tono, tini

	//end variable

	//constant dan operator
	const full_name string = "Muhammad Rizky Alifian Sarodi"
	println(full_name)
}
