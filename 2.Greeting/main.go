package main

import "fmt"

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Hello,", name)

	// fmt.Println("Enter your first and last name sperated by a space: ")
	// fmt.Scanf(" %s %s", &firstName, &lastName)
	// fmt.Println(firstName)

}
