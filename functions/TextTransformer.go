package functions

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Convert word to lower case
*/

// Handles text tranform command with no number
func TransformerNoNum(arr []string, i int, f func(string) string) []string {
	punctuation := strings.Split(arr[i], ")")[1]
	// Convert the previous string to lowercase and append the split function name
	if arr[i] == "(cap)" {
		arr[i-1] = f(string(arr[i-1][0])) + arr[i-1][1:]
	} else {
		arr[i-1] = f(arr[i-1])
	}
	arr[i-1] = arr[i-1] + punctuation
	// Remove the current element from the slice
	arr = append(arr[:i], arr[i+1:]...)
	return arr
}

// Handles text tranform command with number
func TransformerWithNum(arr []string, i int, f func(string) string) []string {
	num, err := strconv.Atoi(strings.Split(arr[i+1], ")")[0])
	if err != nil {
		fmt.Println("Error converting number at:", arr[i]+" "+arr[i+1], " Wrong format.")
		return arr
	}
	punctuation := strings.Split(arr[i+1], ")")[1]
	if arr[i] == "(cap," {
		for j := i - num; j < i && j >= 0; j++ {
			arr[j] = f(string(arr[j][0])) + arr[j][1:]
		}
	} else {
		for j := i - num; j < i && j >= 0; j++ {
			arr[j] = f(arr[j])
		}
	}
	// Append the remaining punctuations
	arr[i-1] = arr[i-1] + punctuation
	// Remove the current element and the next element from the slice
	arr = append(arr[:i], arr[i+2:]...)
	return arr
}

// Converter
func BinToDecimal(s string) string {
	newDecimal, err := strconv.ParseInt(s, 2, 16)
	if err != nil {
		fmt.Println("Error converting bin to decimal")
	}
	return strconv.Itoa(int(newDecimal))
}

// Converter
func HexToDecimal(s string) string {
	newDecimal, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		fmt.Println("Error converting hex to decimal")
	}
	return strconv.Itoa(int(newDecimal))
}

