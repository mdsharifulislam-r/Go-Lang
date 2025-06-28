package main

import "fmt"

func main() {
	var mapsData = make(map[string]string)
	// string map
	mapsData["name"] = "Shariful"
	fmt.Println(mapsData)

	//int map
	var intMap = make(map[string]int)
	intMap["price"] = 32

	// dynamic value map

	var dynamicMap = make(map[string]interface{})

	dynamicMap["name"] = "Sharif"

	dynamicMap["roll"] = 32

	// delete element

	delete(dynamicMap, "roll")

	// clear map

	clear(dynamicMap)

	// check is that ok or not

	value, ok := intMap["roll"]

	fmt.Printf("%s %v", value, ok)

}
