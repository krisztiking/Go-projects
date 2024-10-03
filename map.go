package main

import (
	"fmt"
)

func testMap() {

	var mp map[string]int = map[string]int{
		"apple":  1,
		"orange": 2,
		"lime":   3,
	}
	fmt.Println(mp)
	fmt.Println(mp["orange"])
}
