package main

import "fmt"

type Student struct {
	ID int
	Name string
	GPA float64
}

func main() {
	numberA := 5
	numberB := &numberA // numberB menyimpan alamat memori dari numberA (bukan nilai numberA)

	fmt.Println(numberA)
	fmt.Println(numberB)
	fmt.Println(*numberB) // * untuk mengambil nilai dari alamat memori yang disimpan di numberB

	*numberB = 10 // *numberB merupakan alamat memori dari numberA, sehingga nilai dari numberA bisa diubah

	fmt.Println(numberA)
	fmt.Println("==========================")

	var numberC int = 5
	var numberD *int = &numberC // numberD merupakan alamat memori dari numberC

	fmt.Println(numberC)
	fmt.Println(*numberD)

	numberC = 20

	fmt.Println(numberC)
	fmt.Println(*numberD)
	fmt.Println("==========================")

	
	change(&numberC, 30)

	fmt.Println(numberC)
	fmt.Println(*numberD)
	fmt.Println("==========================")

	student := Student{1, "Komarudin", 3.5}
	fmt.Println(student.Name)
	gradute(&student)
	fmt.Println(student.Name)

}

// penggunaan pointer sebagai parameter
func change(old *int, new int)  {
	*old = new
}

// penggunaan pointer struct sebagai parameter
func gradute(student *Student)  {
	student.Name = student.Name + " S.Kom"
}