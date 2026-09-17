package main

import "fmt"

// how to make a variadic function
// Now this sum() function only takes int type
func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}
	return total
}

// If we want to make a function that take any Data Type then we need to use ...any or ...interface.

func anyType(nums ...any) int {
	fmt.Println(nums...)
	return len(nums)
}
