package main

import (
	"github.com/fatih/color"
	"github.com/mdsharifulislam-r/first/auth"
)

func main() {
	auth.Login("sharif", "1234")
	color.Red("Hello World")
}
