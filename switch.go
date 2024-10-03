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

	y := 10
	switch y {
	case 20:
		fmt.Println(y)
	default:
		fmt.Println("y != 20")
	}

}
