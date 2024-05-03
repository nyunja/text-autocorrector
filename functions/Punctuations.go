package functions

import (
	"regexp"
	"strings"
	"unicode"
)

/*
Moves punctuations to the left
*/

func Punctuations(arr []string) []string {
	var firstPunctuations string
	for i := 0; i < len(arr); i++ {
		// Move punctuation to end of previous slice
		for len(arr[i]) > 0 && unicode.IsPunct(rune(arr[i][0])) {
			if i > 0 {
				arr[i-1] = arr[i-1] + string(arr[i][0])
			}
			if i == 0 {
				firstPunctuations += string(arr[i][0])
			}
			arr[i] = arr[i][1:]
		}
		// Delete empty slices and recurse
		if len(arr[i]) == 0 {
			arr = Spacing(arr)
			arr = Punctuations(arr)
		}
		// Replace deleted punctuations at the beginning then set the firstpunctuation is empty
		arr = Spacing(arr)
		arr[0] = firstPunctuations + arr[0]
		firstPunctuations = ""

	}
	// Fixes the position of the single quotes
	s := strings.Join(arr, " ")
	re := regexp.MustCompile(`'(\s*)([^']+)(\s*)'`)
	s = re.ReplaceAllString(s, " '$2'")

	// Checks and fixes [] {} and () bracket patterns
	re = regexp.MustCompile(`(\(\s*)([^\(]+)\s*\)`)
	s = re.ReplaceAllString(s, " ($2)")
	re = regexp.MustCompile(`(\[\s*)([^\[]+)\s*\]`)
	s = re.ReplaceAllString(s, " [$2]")
	re = regexp.MustCompile(`(\{\s*)([^\{]+)\s*\}`)
	s = re.ReplaceAllString(s, " {$2}")
	
	// Fix spacing and create string slice to return
	arr = strings.Fields(s)
	arr = Spacing(arr)

	return arr
}

// // Fixes the position of the single quotes
// func SingleQuote(arr []string) []string {
// 	s := strings.Join(arr, " ")
// 	re := regexp.MustCompile(`'(\s*)([^']+)(\s*)'`)
// 	s = re.ReplaceAllString(s, " '$2'")
// 	arr = strings.Fields(s)
// 	return arr
// }

// Checks and fixes [] {} and () bracket patterns
// func Bracket(arr []string) []string {
// 	s := strings.Join(arr, " ")
// 	re := regexp.MustCompile(`(\(\s*)([^\(]+)\s*\)`)
// 	s = re.ReplaceAllString(s, " ($2)")
// 	re = regexp.MustCompile(`(\[\s*)([^\[]+)\s*\]`)
// 	s = re.ReplaceAllString(s, " [$2]")
// 	re = regexp.MustCompile(`(\{\s*)([^\{]+)\s*\}`)
// 	s = re.ReplaceAllString(s, " {$2}")
// 	return strings.Fields(s)
// }

// Removes extra spaces in the output
func Spacing(arr []string) []string {
	str := strings.Join(arr, " ")
	return strings.Fields(str)
}
