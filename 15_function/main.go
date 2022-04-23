package main

import "fmt"

func main() {
	printHello()

	result := add(10, 2)
	fmt.Println("10 + 2 = ", result)

	luas, keliling := calculate(10, 2)
	fmt.Println("luas = ", luas)
	fmt.Println("keliling = ", keliling)

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("10 / 0 = ", result)
}

// function dasar
func printHello() {
	fmt.Println("Hello World")
}

// functtion dengan parameter dan return
func add(a, b int) int {
	return a + b
}

// function dengan multiple return
func calculate(panjang, lebar int) (luas, keliling int) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)
	return
}

// function dengan error handling
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("b tidak boleh 0")
	}
	return a / b, nil
}