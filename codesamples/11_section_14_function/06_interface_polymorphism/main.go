package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}
type secretagent struct {
	person
	ltk bool
}
type human interface {
	speak()
}

func (p person) speak() {
	fmt.Println(p.first, p.last, "calling from person")

}

func (s secretagent) speak() {
	fmt.Println(s.first, s.last, "calling from secrate agent")
}

func bar(h human) {
	fmt.Println("I'm passerd bar ", h)
}

func main() {
	sa1 := secretagent{
		person: person{
			first: "Nihad",
			last:  "Hossain",
		},
		ltk: true,
	}
	sa2 := secretagent{
		person: person{
			first: "Kissing",
			last:  "Parkendy",
		},
		ltk: false,
	}
	p1 := person{
		first: "Hanery",
		last:  "Scotiya",
	}
	// fmt.Println(sa1)
	// fmt.Println(sa2)
	// sa1.speak()
	// sa2.speak()
	bar(sa1)
	bar(sa2)
	bar(p1)

}
