package main

import (
	"fmt"
	"os"

	"goreloaded/functions"
)

func main() {
	// Pick the filename arguments and check for errors
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Please include the input and output filenames as arguments")
		return
	}
	inputFileName := args[0]
	outputFileName := args[1]

	// Read input file and check for errors and edit text
	content, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	// Call function to edit input
	content = functions.FinalizedText(content)

	// Write the edited utput to file
	os.WriteFile(outputFileName, content, 0o644)
}
