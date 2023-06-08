package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to Go Lang"
	fmt.Println(welcome)

	// we can take the use input by using the bufio and os binaries/modules.
	reader := bufio.NewReader(os.Stdin) // here we are creating a reader by using a bufio of os.Stdin(standard input)
	fmt.Println("Enter the rating for our Pizza:")

	// comma ok || err ok

	// and here we are reading the input till a new line('\n') found.
	// with the comma ok syntax we can check if there is an error.
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input)
}
