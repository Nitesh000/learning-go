package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")
	fmt.Println("----------------")

	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all the languages: ", languages)
	fmt.Println("JS : ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all the languages: ", languages)

	// loopsa are interesting in golang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
