package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p *person) speak() {
	fmt.Println(p.first, p.last, p.age)

}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()

}

func main() {
	// p := person{
	// 	first: "Nihad",
	// 	last:  "Hossain",
	// 	age:   21,
	// }
	p1 := person{
		first: "MS",
		last:  "Dynia",
		age:   21,
	}
	p2 := person{
		first: "MS",
		last:  "Dynia",
		age:   21,
	}
	//doesn't work
	//saySomething(p)
	saySomething(&p1)
	p2.speak()

}
