package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p *person) speak() string {
	return fmt.Sprintln(p.first, p.last)

}

type human interface {
	speak() string
}

func saySomething(h human) {
	fmt.Println(h.speak())

}

func main() {
	p1 := person{
		first: "Nihad",
		last:  "Hossain",
		age:   21,
	}
	saySomething(&p1)
	s := p1.speak()
	fmt.Println(s)

}
