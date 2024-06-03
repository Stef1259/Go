package main

import (
	"fmt"
	"unicode"
)

func main() {
	var password string

	fmt.Println("Enter your password")
	fmt.Scanln(&password)
	strenghtCheck(password)

}

func strenghtCheck(password string) {
	var strenghtScore int = 0

	if len(password) >= 8 {
		strenghtScore++
	}

	if hasUpper(password) == true {
		strenghtScore++
	}

	if hasLower(password) == true {
		strenghtScore++
	}
	if hasDigit(password) == true {
		strenghtScore++
	}

	fmt.Printf("strenght of the password is %v", calcStrenght(strenghtScore))
	// fmt.Println(strenghtScore)

}

func hasUpper(password string) bool {
	var hasUp bool = false
	for _, char := range password {
		if unicode.IsUpper(char) == true {
			hasUp = true
		}
	}
	// fmt.Println(hasUp)
	return hasUp
}

func hasLower(password string) bool {
	var hasLow bool = false
	for _, char := range password {
		if unicode.IsLower(char) == true {
			hasLow = true
		}
	}
	return hasLow
}

func hasDigit(password string) bool {
	var hasDigt bool = false
	for _, char := range password {
		if unicode.IsDigit(char) == true {
			hasDigt = true
		}
	}
	return hasDigt
}

func calcStrenght(score int) string {
	var strenght string

	if score >= 3 {
		strenght = "strong"

	} else {
		strenght = "weak"
	}

	return strenght

}
