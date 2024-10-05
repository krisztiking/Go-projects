package main

import "fmt"

func changer(str *string) {
	*str = "changed"
}

func changer2(str string) {
	str = "changed twice"
}

func pointers() {
	x := 7
	y := &x
	fmt.Println(x, y)
	*y = 8
	fmt.Println(x, y)

	alf := "hello"
	fmt.Println(alf)
	changer(&alf)
	fmt.Println(alf)

	changer2(alf)
	fmt.Println(alf)
}
