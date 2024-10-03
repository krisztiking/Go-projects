package main

import "fmt"

func main() {

	arr := [5]int{5, 7, 9, 11, 13}
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	fmt.Println(sum)
	fmt.Println(arr)

	arr2D := [2][3]int{{1, 3, 5}, {2, 4, 6}}
	fmt.Println(arr2D)
	fmt.Println(arr2D[1][2])
}
