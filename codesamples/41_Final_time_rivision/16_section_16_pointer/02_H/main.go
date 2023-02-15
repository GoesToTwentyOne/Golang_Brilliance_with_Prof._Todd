package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

func changeme(p *person) {
	(*p).first = "Alex Goot"

}
func main() {
	p := person{
		first: "Md. Nihad",
		last:  "Hossain",
	}
	fmt.Println(p)
	changeme(&p)
	fmt.Println(p)

}
