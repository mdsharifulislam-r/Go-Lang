package main

import "fmt"

// rest opareator
func variadic(nums ...int) int {
	sum := 0

	for _, value := range nums {
		sum += value
	}

	return sum
}

func main() {
	nums := []int{1, 3, 4, 5, 6}
	nums2 := []int{}
	// spread oparator
	var sum int = variadic(nums...)

	nums2 = append(nums2, nums...)

	fmt.Println(nums2)

	fmt.Println(sum)

}
