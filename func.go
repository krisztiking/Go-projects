package main

import "fmt"

func add(x, y int, z string) {
	fmt.Println(x+y, z, "- add")
}

func add2(x int, y int, z string) (int, string, int) {
	return x + y, z, x - y

}

func add3(x, y int, z string) (xx int, zz string) {
	//a functionben a defert a return utan vegzi el
	defer fmt.Println("Done")
	xx = x + y
	zz = z + "bye"
	return
}

func function() {
	//add
	add(1, 2, "szia")

	//add2
	ans1, ans2, ans3 := add2(3, 4, "bye")
	fmt.Println(ans1, ans2, ans3, " - add2")

	//add3
	ans1, ans2 = add3(5, 6, "hi")
	fmt.Println(ans1, ans2)
}
