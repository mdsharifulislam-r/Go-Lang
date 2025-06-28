package main

import "fmt"

func changeNumber(a *int) {
	*a = 4

}

func main() {
	a := 10

	fmt.Println("Number is", a)
	changeNumber(&a)

	fmt.Println("After change", a)
}
