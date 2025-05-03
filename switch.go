package main

import "fmt"

func main() {
	var a int = 10
	switch a {
	case 10:
		fmt.Println("a is equal to 10")
		break
	case 20:
		fmt.Println("a is equal to 20")
		break
	default:
		fmt.Println("default case")
	}

	whoamI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("i am boolean")
		case float64:
			fmt.Println("i am float64")
		case string:
			fmt.Println("i am string")
		default:
			fmt.Printf("%T", i)
		}

	}
	whoamI(5.3)
	whoamI(true)
	whoamI("hello")
}
