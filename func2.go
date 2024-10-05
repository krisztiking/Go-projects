package main

import "fmt"

func tester(x int) {
	fmt.Println("Hello world!", x)
}

func tester2(t2 func(int) int) {
	fmt.Println(t2(12))
}

func returnFunc2(x string) func() {
	
	return func() { fmt.Println(x) }
}

func function2() {
	x := tester
	x(5)

	tester1 := func(x int) int {
		return x * 3
	}
	tester2(tester1)
	// fmt.Println(tester1(11))

	func() {
		fmt.Println("szimpla")
	}()
	returnFunc2("Na eleg")()

}
