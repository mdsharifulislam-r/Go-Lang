package main

import (
	"fmt"
	"os"
)

func main() {
	f, e := os.Open("example.txt")

	if e != nil {
		panic(e)
	}

	fileBuffer := make([]byte, 10)

	file, _ := f.Read(fileBuffer)

	for i := 0; i < len(fileBuffer); i++ {
		fmt.Println(file, string(fileBuffer[i]))
	}

}
