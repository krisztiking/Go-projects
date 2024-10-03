package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("How old are you?")
	scanner.Scan()
	age, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	if age >= 90 {
		fmt.Println("You cannot have a drivers licence, because you are too old.")
	} else if age >= 18 {
		fmt.Println("You can have a drivers licence. \nDo you have?")
	} else {
		fmt.Println("You cannot have a drivers licence")
		fmt.Printf("You have to wait %d years to have a drivers licence.", 18-age)
	}

}
