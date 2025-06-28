package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 12; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < 4 {
		fmt.Println(i)
		i++

	}

	for n := range 6 {
		fmt.Print(n)
	}
}
