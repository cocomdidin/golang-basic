package main

import "fmt"

func main() {
	// map adalah sebuah tipe data yang bisa menampung banyak nilai dengan key yang unik
	// map dapat menggunakan string sebagai key

	myMap := make(map[string]int)

	myMap["key1"] = 1
	myMap["key2"] = 2
	myMap["key3"] = 3

	for i, v := range myMap {
		fmt.Println("myMap ", i, " berisi ", v)
	}


	fmt.Println("===========================")

	bahasa := map[string]string{
		"php": "dynamic",
		"go":  "static",
	}

	bahasa["python"] = "static"

	for i, v := range bahasa {
		fmt.Println("bahasa ", i, " bersifat ", v)
	}


	fmt.Println("===========================")

	delete(bahasa, "go")

	for i, v := range bahasa {
		fmt.Println("bahasa ", i, " bersifat ", v)
	}
}