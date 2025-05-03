package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	fmt.Println(m["a"])
	fmt.Println(len(m))
	delete(m, "a")
	fmt.Println(m)
	v, ok := m["a"]
	fmt.Println(v, ok)
	for k, v := range m {
		fmt.Println(k, v)
	}
}
