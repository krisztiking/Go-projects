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

	var a []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a)
	a = append(a, 10)

	fmt.Println(a)

	var b [6]int = [6]int{11, 22, 33, 44, 55, 66}
	fmt.Println(b)
	fmt.Println(cap(b))
	c := b[0:3]
	fmt.Println(c)
	d := b[3:5]
	fmt.Println(d)
	fmt.Println(cap(d))
	d = append(d, b[5])
	fmt.Println(d)

	e := make([]int, 5)
	fmt.Println(e)

	//array szamok kivalogatasa es sorrendbe rakasa
	var f []int = []int{1, 4, 7, 5, 2, 9, 5, 3, 4, 1, 55, 33, 11, 11, 11}
	hu := make([]int, 0)
	fmt.Println(f)
	for i, element := range f {
		for j, element2 := range f {

			if element == element2 && i != j {
				break
			} else if element != element2 && j == len(f)-1 {
				hu = append(hu, element)
			} else if element == element2 && i == j && j == len(f)-1 {
				hu = append(hu, element)
			} else if element != element2 {
				continue
			}

		}
	}

	fmt.Println(hu)

	//novekvo sorrend

	for i, element := range hu {
		for j := i + 1; j < len(hu); j++ {
			if element > hu[j] {
				vari := hu[i]
				hu[i] = hu[j]
				hu[j] = vari
			}
		}
	}
	fmt.Println(hu)

}
