package main

import (
	"fmt"
)

func groupWordsByFirstWord(arr []string) map[string][]string {
	mapsData := make(map[string][]string)

	for _, v := range arr {
		strV := string(v[0])

		if mapsData[strV] != nil {
			mapsData[strV] = append(mapsData[strV], v)

		} else {
			mapsData[strV] = []string{v}
		}
	}
	return mapsData
}

func main() {
	sliceData := []string{"apple", "ant", "book", "bat"}

	result := groupWordsByFirstWord(sliceData)

	fmt.Println(result)
}
