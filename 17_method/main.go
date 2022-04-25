package main

import "fmt"

type User struct {
	Id       int
	Username string
	Name     string
	Age      int
	Email    string
}

func (user User) display() string {
	return fmt.Sprintf("Nama: %s, Umur: %d, Email: %s", user.Name, user.Age, user.Email)
}

func main() {
	user1 := User{
		Id:       1,
		Username: "user1",
		Name:     "User 1",
		Age:      20,
		Email:    "user1@gmail.com",
	}

	fmt.Println(user1.display())
	// Method sejatinya adalah function namun hanya dimiliki oleh struct
}
