package main

import (
	"fmt"
)

func main() {
	var choice int
	var value float32

	fmt.Println("please enter 1 to conver celcius to farenheit and number 2 to convert farenheit to celcius")
	fmt.Scanln(&choice)

	fmt.Println("please enter the value you want to calculate")
	fmt.Scanln(&value)

	switch choice {
	case 1:
		fmt.Printf("value of %v converted into farenheit is %v", value, convertToFarenheit(value))
	case 2:
		fmt.Printf("value of %v converted into celcius is %v", value, convertToCelcius(value))
	}

}

func convertToFarenheit(value float32) float32 {

	return value*(9.0/5.0) + 32.0
}

func convertToCelcius(value float32) float32 {

	return (value - 32.0) * (5.0 / 9.0)
}
