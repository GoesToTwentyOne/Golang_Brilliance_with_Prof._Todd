package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(p.first)
	fmt.Println(p.age)
}

func main() {
	p := person{
		first: "Nihad",
		last:  "Hossain",
		age:   21,
	}
	p.speak()

}
