package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Method says that Person name : ", p.first, " & age : ", p.age)

}

func main() {
	p1 := person{
		first: "Nihad",
		last:  "Hossain",
		age:   21,
	}
	p2 := person{
		first: "Senya",
		last:  "Goot",
		age:   21,
	}
	p1.speak()
	p2.speak()
}
