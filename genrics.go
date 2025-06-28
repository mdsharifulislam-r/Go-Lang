package main

import "fmt"

func print[T interface{}](item T) {

	fmt.Println(item)

}

func main() {

	print(3)

}
