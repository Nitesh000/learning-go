package main

import "fmt"

func main() {
	fmt.Println("Welcome to golang pointers.")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23
	ptr := &myNumber
	fmt.Println("Value of actual pointer is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("Value after adding 2 is ", myNumber)
}
