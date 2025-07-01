package main

import "fmt"

func sumOfArrayByForLoop(nums ...int) int {
	sum := 0

	for _, val := range nums {
		sum += val
	}

	return sum
}

func main() {

	// var arr = [3]int{1, 3, 4}

	result := sumOfArrayByForLoop(1, 2, 4, 6, 1)

	fmt.Println(result)

}
