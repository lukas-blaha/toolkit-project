package main

import (
	"fmt"

	"github.com/lukas-blaha/toolkit"
)

func main() {
	var tools toolkit.Tools

	s := tools.RandomString(10)
	fmt.Println("Random string:", s)
}
