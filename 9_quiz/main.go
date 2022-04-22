package main

import "fmt"

func main() {

	// menampilkan karakter hanya pada index genap dan huruf vokal dari sebuah string 

	title := "Golang the best language"

	for idx, char := range title {
		charString := string(char)

		if idx % 2 == 0 {
			switch charString {
			case "a", "e", "i", "o", "u":
				fmt.Println("index ", idx, " char ", charString)
			}
		}
			
	}
}