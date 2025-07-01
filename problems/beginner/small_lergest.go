package main

import "fmt"

func secondLargestNumber(nums ...int) int {
	large := 0
	secondLerge := 0
	for _, val := range nums {
		if val > large {
			temp := large
			large = val
			secondLerge = temp
		}

	}
	return secondLerge

}

func main() {

	result := secondLargestNumber(1, 2, 4, 5, 2, 1, 3, 6)

	fmt.Println(result)
}
