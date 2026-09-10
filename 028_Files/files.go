package main

import (
	"fmt"
	"os"
)

func main() {

	// 1. File Information **********
	f, err := os.Open("example.txt")
	// The os.Open() return two things.
	// It return two things.
	// 1. File-Object which is pointer
	// 2. error
