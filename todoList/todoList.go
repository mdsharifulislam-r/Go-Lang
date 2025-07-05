package main

import (
	"fmt"
	"os"
)

type oparationType string

const (
	ADD    oparationType = "add"
	DELETE               = "delete"
	LIST                 = "list"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("=============Please enter 2 argument=================")
	}

	var oparator = os.Args[1]
	switch oparator {
	case "add":

	}

}
