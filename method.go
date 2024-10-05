package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

func (s *Student) setAge(age int) {
	s.age = age
}

func (s Student) getAverageGrade() float32 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

func (s Student) getMaximumGrade() int {
	a := 0
	for i := 0; i < len(s.grades)-1; i++ {
		if s.grades[i] > s.grades[i+1] && s.grades[i] > a {
			a = s.grades[i]
		}

	}
	return a
}

func testMethod() {

	s1 := Student{"Tim", []int{6, 8, 9, 10, 1, 2, 3, 4}, 19}
	fmt.Println(s1)
	s1.setAge(7)
	fmt.Println(s1)

	average := s1.getAverageGrade()
	fmt.Println("avereage grade: ", average)

	max := s1.getMaximumGrade()
	fmt.Println("Max grade: ", max)
}
