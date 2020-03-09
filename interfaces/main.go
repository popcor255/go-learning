package main

import "fmt"

type Person interface {
	// age is a bad example but i cant think of anything else
	getAge() int
}

type student struct {
	name  string
	age   int
	grade int
}

type teacher struct {
	name       string
	age        int
	employeeID int
}

type collegeStudent struct {
	student
}

func (s student) getAge() int {
	return s.age
}

func (t teacher) getAge() int {
	return t.age
}

func whoIsOlder(p1 Person, p2 Person) {
	fmt.Println(p1.getAge(), " is older than ", p2.getAge())
}

func main() {
	tom := student{"tom", 12, 90}
	bill := teacher{"kevin", 19, 122}
	will := collegeStudent(student: student{"tom", 12, 90})
	whoIsOlder(tom, bill)
}
