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

	/*
		Full Struct value {11 Alok alok@gmail.com 25}
		Individual Struct value Only ID 11
		Individual Struct value Only Name Alok
		Individual Struct value Only Email alok@gmail.com
		Individual Struct value Only Age 25
	*/

	// struct fields are mutable bydefault
	// Bsically we can change the value if required.

	u1.Age = 400

	fmt.Println("Struct value after mutation", u1)

	// Output = Struct value after mutation {11 Alok alok@gmail.com 400}

}
