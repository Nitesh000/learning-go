package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")
	fmt.Println("----------------------------")

	days := []string{"Sunday", "Tuesday", "Wednsday", "Friday", "Saturday"}

	fmt.Println(days)

	// for loops in golang
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for _, day := range days {
	// 	fmt.Printf("The day is : %v\n", day)
	// }

	rougueValue := 1
	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco
		}
		if rougueValue == 5 {
			break
		}

		fmt.Println("Value is : ", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jumping at niteshtudu.com")
}
