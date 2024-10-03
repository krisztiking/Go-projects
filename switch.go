package main

import (
	"fmt"
)

func main() {

	x := 10

	switch {
	case x > 1:
		fmt.Println(x)
	default:
		fmt.Println("OK")

	}

}
