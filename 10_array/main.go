package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a)

	b := [2]string{"Hello", "World"}
	fmt.Println(b)

	var c [2][2]string
	c[0][0] = "Hello"
	c[0][1] = "World"
	c[1][0] = "Hai"
	c[1][1] = "Dunia"
	fmt.Println(c)

	d := [...]string{
		"golang",
		"php",
		"java",
	}
	fmt.Println(d)
	fmt.Println("jumlah data ", len(d))

	for idx, val := range d {
		fmt.Println("array index ", idx, " berisi ", val)
	}
}