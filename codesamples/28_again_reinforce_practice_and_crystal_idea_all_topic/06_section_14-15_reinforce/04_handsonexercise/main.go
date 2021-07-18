package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("Name: %s Age: %d \n", p.first, p.age)

}
func main() {
	per := person{
		first: "Nihad",
		last:  "Hossain",
		age:   21,
	}
	per.speak()

}
