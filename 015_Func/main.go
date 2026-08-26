package main

import "fmt"

/*
In this line "func add(a int, b int) int {"

this line represent the input taken by function add(a int, b int) and type is int.

but this int { define that there is a return type and must be int.

*/

func add(a int, b int) int {
	return a + b
}

func main() {
	result := add(3, 5)
	fmt.Println(result)
}
