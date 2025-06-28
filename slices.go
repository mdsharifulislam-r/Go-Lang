package main

import (
	"fmt"
	"slices"
)

func main() {
	a := [3]string{"a", "b"}

	b := make([]string, len(a))
	copy(b, a)

	fmt.Print(b)

}
