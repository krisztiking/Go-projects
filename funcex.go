package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func a2(x int64, i int64, y bool) bool {
	if x%i == 0 {
		y = false
		fmt.Println("nem prim")
		return y
	} else {
		//fmt.Println("eddig prim")
		return y
	}
}

func funcEx() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	szm, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	ih := true
	var i int64 = 0

	for i = 2; i <= 7; i++ {

		if a2(szm, i, ih) == false {
			ih = false
			break
		}
	}

	if ih == true {
		fmt.Println("prim")
	} else {
		fmt.Println("nem prim")
	}
	fmt.Println(a2)
}
