package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	fmt.Println("before dereferencing", p.first)
	fmt.Println("before dereferencing", p.last)
	fmt.Println("before dereferencing", p.age)
	(*p).first = "Nihad"
	//another way
	p.last = "Hossain"
	(*p).age = 21
	fmt.Println("after dereferencing", p.first)
	fmt.Println("after dereferencing", p.last)
	fmt.Println("after dereferencing", p.age)

}

func main() {
	p := person{
		first: "Aneyaliya",
		last:  "Goot",
		age:   27,
	}
	fmt.Println("before referencing ", p)
	changeMe(&p)
	fmt.Println("after referencing ", p)

}
