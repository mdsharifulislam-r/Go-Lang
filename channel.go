package main

import "fmt"

func sum(chanel chan int) {

	chanel <- 5

}
func main() {
	chanData := make(chan int)

	go sum(chanData)

	number := <-chanData

	fmt.Println((number))
}
