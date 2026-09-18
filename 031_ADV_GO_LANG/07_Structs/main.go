package main
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
