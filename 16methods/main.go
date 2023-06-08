package main

import "fmt"

func main() {
	fmt.Println("Methods in golang")
	fmt.Println("---------------------")
	// no inheritance in golang; No super or no parent

	nitesh := User{"Nitesh", "ntudu040@gmail.com", true, 22}
	fmt.Println(nitesh)
	// another way of format, with this it also shows the key and value
	// The plus is use for the both key and value representation
	fmt.Printf("Nitesh's details are: %+v\n", nitesh)
	fmt.Printf("Name is %v and email is %v.\n", nitesh.Name, nitesh.Email)

	// using methods in golang
	nitesh.GetStatus()
	nitesh.NewMail()
	fmt.Printf("Name is %v and email is %v.\n", nitesh.Name, nitesh.Email)
}

// The structure names should be of Pascal casing
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	// we can also create some private fields by using the camel casing
	// for example we can create a private field for password
	// password string
}

// way to define methods for the structs
// methods visibilityis also the same as the field in struct for camel casing and pascal casing

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	// by doing this we can't update the values because it passes a copy of the value. But not the reference.
	// So we have to use the pointers
	// that is a defer statement. and pointer concepts.
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is : ", u.Email)
}
