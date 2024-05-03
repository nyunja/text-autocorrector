package functions

import (
	"strings"
)

/*
CommandChecker processes a slice of strings based on specific function markers.
*/

func CommandChecker(arr []string) []string {
	// Iterate through each element in the input slice
	for i := 0; i < len(arr); i++ {
		if strings.Contains(arr[i], "(cap)") {
			arr = TransformerNoNum(arr, i, strings.ToUpper)
		} else if arr[i] == "(cap," {
			arr = TransformerWithNum(arr, i, strings.ToUpper)
		} else if strings.Contains(arr[i], "(low)") {
			arr = TransformerNoNum(arr, i, strings.ToLower)
		} else if arr[i] == "(low," {
			arr = TransformerWithNum(arr, i, strings.ToLower)
		} else if strings.Contains(arr[i], "(up)") {
			arr = TransformerNoNum(arr, i, strings.ToUpper)
		} else if arr[i] == "(up," {
			arr = TransformerWithNum(arr, i, strings.ToUpper)
		} else if strings.Contains(arr[i], "(bin)") {
			arr = TransformerNoNum(arr, i, BinToDecimal)
		} else if arr[i] == "(bin," {
			arr = TransformerWithNum(arr, i, BinToDecimal)
		} else if strings.Contains(arr[i], "(hex)") {
			arr = TransformerNoNum(arr, i, HexToDecimal)
		} else if arr[i] == "(hex," {
			arr = TransformerWithNum(arr, i, HexToDecimal)
		}
	}
	return arr
}
