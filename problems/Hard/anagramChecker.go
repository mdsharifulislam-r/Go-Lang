package main

import "fmt"

func isAnagram(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	firstCount := 0
	secondCount := 0

	for i := 0; i < len(str1); i++ {
		firstCount += int(str1[i])
		secondCount += int(str2[i])
	}

	return firstCount == secondCount

}

func main() {

	result := isAnagram("and", "dna")

	fmt.Println(result)
}
