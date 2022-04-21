package main

import (
	"fmt"
	"package/calculation"
)

func main() {
	fmt.Println("Hello, World!")

	sentence := TestAja()
	fmt.Println(sentence)

	result := calculation.Add(10, 5)
	fmt.Println(result)
}

/*

1. Package executable
	aturan golang package executable harus bernama main
2. Package library
	package selain main

*/