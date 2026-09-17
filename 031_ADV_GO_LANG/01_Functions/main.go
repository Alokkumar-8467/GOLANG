package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func SumAndProduct(a int, b int) (int, int) {
	sum := a + b
	product := a * b

	return sum, product

}
