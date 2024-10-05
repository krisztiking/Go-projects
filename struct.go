package main

import "fmt"

type Point struct {
	x int32
	y int32
}
type Circle struct {
	radius float64
	center *Point
	//*Point
}

func changeX(c *Circle) {
	c.center.x = 0
}

func structure() {
	c1 := &Circle{3.45, &Point{4, 5}}
	fmt.Println(c1)
	// cc := *c2
	// fmt.Println(cc)
	changeX(c1)
	fmt.Println(c1)

	/*var p1 Point = Point{1, 2, false}
	//	p1 := Point{1, 2, false}
	var p2 Point = Point{-5, 7, true}
	fmt.Println("p1 = ", p1.x, p1.y)
	fmt.Println("p2 = ", p2.x, p2.y)

	p1.x = 12
	fmt.Println("p1 new = ", p1.x, p2.z)

	p1 = Point{x: 3, y: 99}
	fmt.Println(p1)*/
}
