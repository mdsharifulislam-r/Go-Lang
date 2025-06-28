package main

import "fmt"

func main() {
	str := "Shariful"

	var revStr []byte

	for i := len(str) - 1; i > -1; i-- {
		revStr = append(revStr, str[i])
	}

	fmt.Println(string(revStr))
}
