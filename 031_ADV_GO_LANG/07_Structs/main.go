package main

import "fmt"

// struct groups related fields into one type

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func main() {

	u1 := User{
		ID:    11,
		Name:  "Alok",
		Email: "alok@gmail.com",
		Age:   25,
	}

	fmt.Println("Full Struct value", u1)
	fmt.Println("Individual Struct value Only ID", u1.ID)
	fmt.Println("Individual Struct value Only Name", u1.Name)
	fmt.Println("Individual Struct value Only Email", u1.Email)
	fmt.Println("Individual Struct value Only Age", u1.Age)

}
