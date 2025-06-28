package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func multipleReturn() (string, int) {
	return "I", 3
}

func funcExpected(add func(a int, b int) int) int {
	return add(2, 4)
}

func main() {
	// single return function
	c := add(3, 5)
	fmt.Println(c)

	//muliple return function
	a, b := multipleReturn()
	fmt.Println(a, b)

	//function expect

	result := funcExpected(func(a, b int) int {
		return a * b
	})

	fmt.Println(result)

}
