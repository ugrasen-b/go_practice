package main

import (
	"fmt"
)

type Student struct {
	name string
	gpa  float32
}

func main() {
	fmt.Println("Using a simple struct")
	var s Student
	ptrTos := &s
	s.name = "Ug"
	ptrTos.gpa = 8.67
	fmt.Println("student is ", s)
	fmt.Println("Student is referenced using pointer :", *ptrTos)
}
