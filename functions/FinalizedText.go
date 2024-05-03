package functions

import (
	"strings"
)

/*
Executes all text commands to edit punctuations and returns final text as a slice of bytes
*/

func FinalizedText(content []byte) []byte {
	arr := strings.Fields(string(content))
	arr = Punctuations(arr)
	//arr = SingleQuote(arr)
	//arr = Bracket(arr)
	arr = Vowels(arr)
	arr = CommandChecker(arr)
	arr = Punctuations(arr)
	//arr = Bracket(arr)
	//arr = SingleQuote(arr)
	arr = Vowels(arr)
	finalContent := strings.Join(arr, " ") + "\n"
	content = []byte(finalContent)
	return content
}
