package main

import "fmt"

func main() {
	score := 90
	var grade string

	if score <= 50 {
		grade = "E"
	} else if score <= 60 {
		grade = "D"
	} else if score <= 70 {
		grade = "C"
	} else if score <= 80 {
		grade = "B"
	} else {
		grade = "A"
	}

	fmt.Println("Hello, Nilai anda", score, "masuk dalam grade", grade)

}