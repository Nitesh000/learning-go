package main

import "fmt"

const LoginToken string = "gibrish"

func main() {
	// The format of writing variabes is like {var/const variable_name type_of_variable}
	// var will be like the variable whcih can be changed.and const wich can't be
	var username string = "Nitesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggidIn bool = true
	fmt.Println(isLoggidIn)
	fmt.Printf("Variable is of type: %T \n", isLoggidIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallfloat float64 = 255.23425253253
	fmt.Println(smallfloat)
	fmt.Printf("Variable is of type: %T \n", smallfloat)

	// Default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	// This is another type of variable defining method in golang.
	// this operator is called (:=) short variable declatation operator
	website := "niteshtudu.com"
	fmt.Println(website)

	// no var style
	numberOfUsers := 400000
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
