package main

import "fmt"

type User struct {
	Id       int
	Username string
	Name     string
	Age      int
	Email    string
}

func main() {
	// Assign value ke struct
	// cara 1
	user1 := User{}
	user1.Id = 1
	user1.Username = "user1"
	user1.Name = "User 1"
	user1.Age = 20
	user1.Email = "user1@gmail.com"

	// cara 2
	user2 := User{
		Id:       2,
		Username: "user2",
		Name:     "User 2",
		Age:      21,
		Email:    "user2@gmail.com",
	}

	// cara 3
	user3 := User{3, "user3", "User 3", 22, "user3@gmail.com"}

	displayUser(user1)
	displayUser(user2)
	displayUser(user3)
}

func displayUser(user User) {
	result := fmt.Sprintf("Nama: %s, Umur: %d, Email: %s", user.Name, user.Age, user.Email)
	fmt.Println(result)
}