package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	fmt.Println("----------------------")
	greet()

	// result := adder(3, 5)

	// fmt.Println("Result is : ", result)

	proResult, messagge := ProAdder(3, 5, 2, 52, 5)
	fmt.Println("Pro result is : ", proResult)
	fmt.Println("Pro result message is : ", messagge)
}

// NOTE: It isn't allow to create a function inside a function.
// NOTE: The order of funciton and funciton doesn't need. We can create them anywhere in the same scope.

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func ProAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "Hii from the ProAdder"
}

func greet() {
	fmt.Println("Namaste! from golang")
}
