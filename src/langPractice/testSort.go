package langPractice

import (
	"fmt"
	"sort"
)

type Student struct {
	age  int
	name string
}

type Students []Student

func (a Students) Len() int {
	return len(a)
}
func (a Students) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a Students) Less(i, j int) bool {
	return a[i].age <= a[j].age
}

func TestSort() {
	s1 := Student{
		age:  14,
		name: "Alice",
	}
	s2 := Student{
		age:  16,
		name: "Bob",
	}
	s3 := Student{
		age:  13,
		name: "Mark",
	}
	students := []Student{s1, s2, s3}
	sort.Sort(Students(students))
	fmt.Println(students)
}
