package main

import "fmt"

func increament() func() int {
	count := 0
	return func() int {
		count += 1

		return count
	}
}

func main() {

	counter := increament()

	fmt.Println(counter())
	fmt.Println(counter())
}
