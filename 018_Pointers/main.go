package main

import "fmt"

// We pass it by value
func changeNum(num int) {
	num = 5
	fmt.Println("In changeNum", num)
}

func main() {

	num := 1

	changeNum(num)
	fmt.Println("After changeNum in main", num)

}

