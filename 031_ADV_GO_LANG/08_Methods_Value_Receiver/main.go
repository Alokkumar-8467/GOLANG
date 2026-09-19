package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {

	u := User{Name: "Alok", Age: 24}
	fmt.Println(u.Intro())

}

func (u User) Intro() string {
	return fmt.Sprintf("Hi, I am %s", u.Name)
}

/*
You're mostly right, with two small corrections.

What's right
Intro is a method.
User is the struct type.
The user is handed to the method, so Intro can use its data.
Correction 1: the receiver is not a normal input
