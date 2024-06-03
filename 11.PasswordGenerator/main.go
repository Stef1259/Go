package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	var hasUp, hasLow, hasDig, password, lenght string

	fmt.Println("select the optins you want for password generation")

	fmt.Println("type in you password lenght")
	fmt.Scanln(&lenght)

	fmt.Println("if you want an uppercase letter type y, if not type n")
	fmt.Scanln(&hasUp)

	fmt.Println("if you  want an lowercase letter type y, if not type n ")
	fmt.Scanln(&hasLow)

	fmt.Println("if you  want a digit type y, if not type n ")
	fmt.Scanln(&hasDig)

	password = passwordGenerator(hasUp, hasLow, hasDig, lenght)
	fmt.Printf("generated password id %v", password)

}

func passwordGenerator(hasUp, hasLow, hasDig, lenght string) string {

	var password string

	passwordlenght, err := strconv.Atoi(lenght)
	if err != nil {
		fmt.Println("invalid password lenght")
	}

	const (
		lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
		uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		digits           = "0123456789"
		allChars         = lowercaseLetters + uppercaseLetters + digits
	)

	for len(password) < passwordlenght {
		if hasUp == "y" {
			randNo := rand.Intn(26)
			newUp := string(uppercaseLetters[randNo])
			password += newUp
		}
		if hasLow == "y" {
			randNo := rand.Intn(26)
			newLow := string(lowercaseLetters[randNo])
			password += newLow
		}

		if hasDig == "y" {
			randNo := rand.Intn(10)
			newDig := string(digits[randNo])
			password += newDig
		}

	}

	return password

}
