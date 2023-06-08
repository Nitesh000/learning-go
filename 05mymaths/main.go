package main

import (
	"fmt"
	// "math/big"
	"math/rand"
	// "crypto/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	// var mynumberOne int = 2
	// var mynumberTwo float64 = 4

	// fmt.Println("The sum is: ", mynumberOne+int(mynumberTwo))

	// random number
	// we need a Seed, because to generate random number the function is depend on a Seed.
	// And the seed should always chagne. so we use the time.Now().UnixMicro() to get the current time in nano.
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(5) + 1)

	// ranadom number from crypto

	// myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	// fmt.Println(myRandomNum)
}
