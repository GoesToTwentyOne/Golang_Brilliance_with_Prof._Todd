package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("he says my name is %v my age is %v \n", p.first, p.age)
}

func main() {
	p1 := person{
		first: "Nihad",
		last:  "Hossain",
		age:   21,
	}
	p1.speak()

}
