package functions

/*
Checks if a word starts with a vowel and the word before it is a or A then changes it to A or An
*/

func Vowels(arr []string) []string {
	vowels := "aeiouhAEIOUH"
	for _, vowel := range vowels {
		for i := 0; i < len(arr); i++ {
			if i > 0 && arr[i][0] == byte(vowel) {
				if arr[i-1] == "a" {
					arr[i-1] = "an"
				}
				if arr[i-1] == "A" {
					arr[i-1] = "An"
				}
			}
		}
	}
	return arr
}
