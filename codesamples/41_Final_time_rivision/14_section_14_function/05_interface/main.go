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
type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Printf(p.first, p.last)
}
func (s secretAgent) speak() {
	fmt.Print(s.person, s.ltk)
}
func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println(h.(person).first)
	case secretAgent:
		fmt.Println(h.(secretAgent).ltk)
	}
}
func main() {
	s1 := secretAgent{
		person: person{
			first: "Nihad",
			last:  "Hossain",
		},
		ltk: true,
	}
	s2 := secretAgent{
		person: person{
			first: "Hena",
			last:  "Pike",
		},
		ltk: false,
	}
	p := person{
		first: "Ira",
		last:  "Hasan",
	}
	bar(s1)
	bar(s2)
	bar(p)

}
