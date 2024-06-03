package main

import (
	"fmt"
)

func main() {

	var choice int
	var num1, num2 float64

	fmt.Println("enter 1 for addition, 2 for substraction, 3 for multiplication, 4 for division")
	fmt.Scanln(&choice)
	fmt.Println("enter the first number")
	fmt.Scanln(&num1)
	fmt.Println("enter the second number")
	fmt.Scanln(&num2)

	switch choice {
	case 1:
		fmt.Println(add(num1, num2))
	case 2:
		fmt.Println(substract(num1, num2))
	case 3:
		fmt.Println(multiply(num1, num2))
	case 4:
		fmt.Println(divide(num1, num2))
	}

}

// function declaration

func add(x, y float64) float64 {
	return x + y

}

func substract(x, y float64) float64 {
	return x - y

}

func multiply(x, y float64) float64 {
	return x * y

}

func divide(x, y float64) float64 {

	if y == 0 {
		fmt.Println("dividing by zero is not allowed.")
	}
	return x / y

}
