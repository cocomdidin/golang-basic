package main

import "fmt"

func main() {
	a := []map[string]string{
		{"nama": "Andi", "umur": "20"},
		{"nama": "Budi", "umur": "30"},
		{"nama": "Caca", "umur": "40"},
	}

	for i, v := range a {
		fmt.Println("a ", i, " berisi ", v)
	}

	fmt.Println("===========================")

	for _, v := range a {
		fmt.Println(v["nama"], " berumur ", v["umur"])
	}
}