package main

import "fmt"

func main() {
	fmt.Println("Belajar variable dan tipe data!")

	// Type data pada golang bersifat static 
	// variabel dengan tipe string tidak bisa di isi dengan integer
	var nama string = "Komarudin"
	var umur int = 30
	var nikah bool = true

	// variable dapat otomatis menentukan tipe data berdasarkan nilai yang di isi
	// tapi tidak dapat mengubah tipe data
	anak := 1

	// default value dari variable adalah sesuai tipe data jika string maka akan mengisi dengan string kosong
	var kosong string

	fmt.Println(kosong)
	fmt.Println("Nama saya", nama)
	fmt.Println("Umur saya", umur)
	fmt.Println("Status nikah saya", nikah)
	fmt.Println("Jumlah anak saya", anak)
}