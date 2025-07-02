package main

import (
	"fmt"
	"strconv"
)

func Map[T any](arr []T, cb func(value T) interface{}) []interface{} {

	temArr := make([]interface{}, len(arr))

	for i := 0; i < len(arr); i++ {
		temArr[i] = cb(arr[i])
	}

	return temArr

}

func main() {
	arr := []int{1, 3, 5}

	result := Map[int](arr, func(value int) interface{} {
		return strconv.Itoa(value)
	})

	fmt.Println(result)
}
