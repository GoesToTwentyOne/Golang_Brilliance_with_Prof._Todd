package main

import "fmt"

type person struct {
	first string
	last  string
}

func changeme(p *person) {
	(*p).first = "Alex"
	(*p).last = "Goot"
}

func main() {
	p := person{
		first: "Nihad",
		last:  "Hossain",
	}
	fmt.Println(p.first)
	fmt.Println(p.last)
	changeme(&p)
	fmt.Println(p.first)
	fmt.Println(p.last)

}
