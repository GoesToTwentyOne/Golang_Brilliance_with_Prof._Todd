package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func changeme(p *person) {
	fmt.Println(p)
	fmt.Println((*p).first)
	(*p).first = "Seneya"
	fmt.Println(p)
	fmt.Println((*p).first)

}

func main() {
	p := person{
		first: "Nihad",
		last:  "Hossain",
		age:   21,
	}
	fmt.Println(p)
	changeme(&p)
	fmt.Println(p)

}
