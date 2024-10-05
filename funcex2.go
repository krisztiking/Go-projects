package main

import "fmt"

func test2(myFunc func(int) int) {
	x := myFunc(7)
	fmt.Println(x)
}

func funcEx2() {
	test := func(x int) int {
		return x * 2
	}
	test2(test)
	fmt.Println(test(2))
}
