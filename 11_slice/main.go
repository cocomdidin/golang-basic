package main

import "fmt"

func main() {

	// slice mirip dengan array, namun slice dapat diisi dengan jumlah data dinamis
	// jadi kita bisa menambahkan data pada slice dengan menggunakan append

	// contoh 1
	var a []string
	a = append(a, "Hello")
	a = append(a, "World")
	fmt.Println(a)

	// contoh 2
	b := []string{"Hello", "World"}
	fmt.Println(b)

	// contoh 3
	c := []string{
		"golang",
		"php",
		"java",
	}
	
	for _, v := range c {
		fmt.Println(v)
	}

}