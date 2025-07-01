package main

import (
	"fmt"
	"strings"
)

func countVowelsFromString(str string) int {
	vowels := map[string]string{"a": "a", "i": "i", "e": "e", "o": "o", "u": "u"}

	count := 0
	lowerStr := strings.ToLower(str)

	for _, v := range lowerStr {
		if vowels[string(v)] == string(v) {
			count += 1
		}
	}

	return count

}

func main() {

	result := countVowelsFromString("MD Shariful Islam")

	fmt.Println(result)
}
