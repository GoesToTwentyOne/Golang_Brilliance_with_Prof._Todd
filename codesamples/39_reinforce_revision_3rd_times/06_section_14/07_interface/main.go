package main

import (
	"fmt"
)

type human interface {
	speak()
}

type person struct {
	first string
	last  string
}
type student struct {
	person
	class string
}

func (s student) speak() {
	fmt.Printf("student information %s %s   \nclass information %s\n", s.first, s.last, s.class)
}
func (p person) speak() {
	fmt.Printf("student information %s %s   \n", p.first, p.last)

}
func information(h human) {
	switch h.(type) {
	case person:
		fmt.Printf("Person %v %v\n", h.(person).first, h.(person).last)
	case student:
		fmt.Printf("Student %v %v %v\n", h.(student).first, h.(student).last, h.(student))
	}

	fmt.Println(h)
}

func main() {
	s1 := student{
		person: person{
			first: "Nihad",
			last:  "Hossain",
		},
		class: "seven",
	}
	s2 := person{
		first: "Jony",
		last:  "Yes",
	}
	information(s1)
	information(s2)

}
