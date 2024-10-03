package main

import (
	"fmt"
)

func main() {

	/*scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("How old are you?")
	scanner.Scan()
	age, _ := strconv.ParseInt(scanner.Text(), 10, 64)*/

	var prim bool = true
	// var om bool = true

	for x := 1; x <= 100; x++ {

		for y := 2; y <= 7; y++ {

			if x%y == 0 {

				if x == y {
					continue
				}
				prim = false
				break
			}

		}
		if prim {
			fmt.Printf("%d \n", x)
		}
		prim = true
	}

}
