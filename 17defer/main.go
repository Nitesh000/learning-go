package main

import "fmt"

// defer: By writing this keyword, what it does is just put the line before the ending of the scope.
// the sequence of defer use like [LIFO] order last in first out.
// so the last line on defer will print first and first one later.

func main() {
	fmt.Println("Welcome to defer in golang")
	fmt.Println("------------------------------")

	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello") // this block will print Hello\n Two\n One\n World\n
	myDefer()            // Hello\n 43210\n Two\n One\n World\n
	// Here the myDefer statement will run first because it is not using the defer here. So after printing hello it will print the funciton.
	// then the other.
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
