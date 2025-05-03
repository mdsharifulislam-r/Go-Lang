package main

import (
	"fmt"
)

func main() {
	var s []string
	fmt.Println(s) // nil
	s = make([]string, 3)
	fmt.Println(s)
	s = append(s, "a", "b")
	fmt.Println(s)
	for i := range s {
		fmt.Printf("%d %v\n", i, s[i])
	}
	fmt.Println(cap(s))
}
