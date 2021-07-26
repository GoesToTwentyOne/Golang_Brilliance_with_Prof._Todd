package main

import "fmt"

type person struct {
	name        string
	age         int
	nationality string
}

func (p *person) speak() string {
	return fmt.Sprintln(p.name, p.age, p.nationality)

}

type human interface {
	speak() string
}

func saySomething(h human) {
	fmt.Println(h.speak())

}

func main() {
	p1 := person{
		name:        "Nihad ",
		age:         21,
		nationality: "Bangladeshi",
	}
	p2 := person{
		name:        "Angela ",
		age:         22,
		nationality: "Germany",
	}
	p3 := person{
		name:        "Ntreyko ",
		age:         27,
		nationality: "Canada",
	}
	fmt.Println(p1, p2, p3)
	//code doesn't run
	//saySomething(p1)
	saySomething(&p1)
	//method sets determines what type implements the interface
	fmt.Println(p1.speak())

}
