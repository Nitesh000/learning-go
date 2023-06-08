package main

import (
	"fmt"
	// "sort"
)

func main() {
	fmt.Println("Welcome to slices tutorial!")
	fmt.Println("----------------------------")

	fruitlist := []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T \n", fruitlist)

	fruitlist = append(fruitlist, "Mango", "Banana")
	fmt.Println("Fruitlist after adding new fruits: ", fruitlist)

	fruitlist = append(fruitlist[:3])
	fmt.Println("Fruitlist after slicing: ", fruitlist)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867
	// highScore[4] = 777

	// highScore = append(highScore, 555, 666, 321)

	fmt.Println(highScore)

	// fmt.Println(sort.IntsAreSorted(highScore))
	// sort.Ints(highScore)
	//
	// fmt.Println(highScore)

	// how to remove a value from slices based on index
	courses := []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
