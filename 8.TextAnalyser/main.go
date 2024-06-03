package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var text, substring, choice string

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter your string ")
	text, _ = reader.ReadString('\n')
	text = strings.TrimSpace(text)
	fmt.Println("press 1 to count the number of words,press 2 to get number of characters and 3 to count the occurances of specific phrases of characters")
	fmt.Scanln(&choice)
	switch choice {
	case "1":
		fmt.Printf("number of words is %v", numWords(text))
	case "2":
		fmt.Printf("number od characters is %v", charCount(text))
	case "3":
		fmt.Println("enter the substring you want to check for")
		substring, _ = reader.ReadString('\n')
		substring = strings.TrimSpace(substring)
		fmt.Printf("number of occurance of substring %v is %v", substring, strings.Count(text, substring))
	}
}

func numWords(text string) int {
	var count int = 1

	for i := range text {

		if text[i] == 32 {
			count++
		}

	}

	return count

}

func charCount(text string) int {
	var count int
	count = len(text)
	return count
}
