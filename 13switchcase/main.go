package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang")
	fmt.Println("---------------------------")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of dice is : ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("Dice value is 2 and youc an move 2 spots")
	case 3:
		fmt.Println("Dice value is 3 and youc an move 3 spots")
		fallthrough
	case 4:
		fmt.Println("Dice value is 4 and youc an move 4 spots")
		fallthrough
	case 5:
		fmt.Println("Dice value is 5 and youc an move 5 spots")
	case 6:
		fmt.Println("Dice value is 6 and youc an move 6 spots")
	default:
		fmt.Println("What was that!!!")
	}
}
