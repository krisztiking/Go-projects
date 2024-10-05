package main

import "fmt"

func returnFunc(x string) func() {
	return func() {
		fmt.Println(x)
	}
}

func funcEx3() {
	returnFunc("Hello")()
	/*x := */ returnFunc("goodbye")()
	//x()
}
