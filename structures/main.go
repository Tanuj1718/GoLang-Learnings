package main

import "fmt"

type Goat struct {
	FirstName string
	LastName  string
	Age       int
}

type User struct{
	Name string
	Age int
	Email string
	BirthYear int
}

func main() {
	var king Goat
	//1st method to fill variable
	king.FirstName = "Virat"
	king.LastName = "Kohli"
	king.Age = 36
	fmt.Println(king.FirstName, king.LastName, king.Age)
	
	//2nd method to fill variable
	hitman:= Goat{
		FirstName: "Rohit",
		LastName: "Sharma",
		Age: 38,
	}
	fmt.Println(hitman)

	Girl := User{
		Name: "Miss PM",
		Age: 20,
		Email: "busy@gmail.com",
		BirthYear: 2004,
	}
	fmt.Println(Girl)
}