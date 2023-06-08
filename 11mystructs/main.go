package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	fmt.Println("---------------------")
	// no inheritance in golang; No super or no parent

	nitesh := User{"Nitesh", "ntudu040@gmail.com", true, 22}
	fmt.Println(nitesh)
	// another way of format, with this it also shows the key and value
	// The plus is use for the both key and value representation
	fmt.Printf("Nitesh's details are: %+v\n", nitesh)
	fmt.Printf("Name is %v and email is %v.", nitesh.Name, nitesh.Email)
}

// The structure names should be of Pascal casing
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
