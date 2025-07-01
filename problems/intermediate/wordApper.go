package main

import (
	"fmt"
	"strings"
)

func wordApperInAString(str string) map[string]int {
	hasmap := make(map[string]int)

	splitText := strings.Split(str, " ")

	for _, val := range splitText {
		if hasmap[val] > 0 {

			hasmap[val] = hasmap[val] + 1

		} else {
			hasmap[val] = 1
		}

	}

	return hasmap
}

func main() {
	result := wordApperInAString("I am a am I I Here")
	fmt.Println(result)
}
