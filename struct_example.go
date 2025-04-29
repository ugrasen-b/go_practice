package main

import (
	"fmt"
)

type Name struct {
	first string
	last  string
}

type Person struct {
	who     Name
	address string
	age     int
}

type Student struct {
	per          Person
	testScores   []int
	gradeAverage int
}

func (s *Student) testAverage() {
	sum := 0
	for _, v := range s.testScores {
		sum += v
	}
	s.gradeAverage = sum / len(s.testScores)
}

func inputScores(s *Student, testCount int) {
	var score int
	fmt.Printf("Input %d scores for %s: ", testCount, s.per.who)
	for i := 0; i < testCount; i++ {
		fmt.Scanf("%d", &score)
		s.testScores[i] = score
	}

}
