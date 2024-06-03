package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var randomNum int = rand.Intn(100)
	var guessingNum int

	fmt.Println(randomNum)
	fmt.Println("please enter your guess for the random number between zero to 100")
	fmt.Scanln(&guessingNum)

	for guessingNum != randomNum {

		fmt.Println("please enter your guess for the random number")
		fmt.Scanln(&guessingNum)

	}
	fmt.Println("You are correct")

}
