package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// in general whatever we take it as an input, it becomes a string.
// so for usecases like taking an integer will be a probalamatic.
// we just need to convert that to integer by using strconv.ParseFloat(string.TrimSpace(input))

func main() {
	fmt.Println("Welcome to our pizza app!!")
	fmt.Println("Please rate out pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	// the ParseFloat will convert it to a floating point number. And the TrimSpace will remove any white spaces incuding the newline('\n').

	// when something is 'none', in golang we use 'nil' as 'none'.
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added to your rating: ", numRating+1)
	}
}
