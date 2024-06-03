package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	var inputFile, outputFile, targetWord string
	var err error

	err = fileProcessor(inputFile, outputFile, targetWord)

	if err != nil {
		fmt.Println("error occured while processing")

	}

}

func fileProcessor(inputFile, outputFilePath, targetWord string) error {

	data, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("error occured in file reading")
	}

	defer data.Close()

	fileContent, err := io.ReadAll(data)
	if err != nil {
		fmt.Println("error occured in reading file content")
	}

	lines := strings.Split(string(fileContent), "\n")
	lineCount := len(lines)
	wordCount := 0
	for _, line := range lines {
		wordCount += strings.Count(line, targetWord)
	}

	output := fmt.Sprintf("Lines: %d\nOccurrences of '%s': %d\n", lineCount, targetWord, wordCount)

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("Failed creating file:", err)
		os.Exit(1)
	}

	// Write the output string to the file
	_, err = outputFile.WriteString(output)
	if err != nil {
		fmt.Println("Failed writing to file:", err)
		os.Exit(1)
	}

	// Close the file
	err = outputFile.Close()
	if err != nil {
		fmt.Println("Failed closing file:", err)
		os.Exit(1)
	}

	return err
}
