package main

import (
	"fmt"
)

func main() {
	var a [5]int
	a[3] = 5
	fmt.Printf("a:%v", a)

	b := [5]int{1, 2, 4, 5}
	fmt.Println(b)

	c := [...]int{1, 2, 4, 5}

	fmt.Println(c)
}
