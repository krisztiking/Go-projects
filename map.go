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
	mp["apple"] = 5

	fmt.Println(mp["apple"])
	fmt.Println(mp)
	delete(mp, "lime")
	fmt.Println(mp)

	// val - felveszi az erteket, ok - true of false
	val, ok := mp["apple"]
	fmt.Println(val, ok)

	val2, ok2 := mp["orange"]
	fmt.Println(val2, ok2)

	// uj ertek hozzaadasa
	mp["banana"] = 13
	fmt.Println(mp)

}
