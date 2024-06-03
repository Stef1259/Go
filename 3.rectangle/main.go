package main

import "fmt"

func main() {

	var lenght, width, area float32

	fmt.Print("Enter the lenght of the rectangle: ")
	fmt.Scanln(&lenght)
	fmt.Print("Enter the width of the rectangle: ")
	fmt.Scanln(&width)

	area = lenght * width

	fmt.Println("The area of the rectangle is: ", area)

}
