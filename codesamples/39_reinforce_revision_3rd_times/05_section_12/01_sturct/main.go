package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Md.Nihad ",
		last:  "Hossain",
		age:   21,
	}
	p2 := person{
		first: "Alex ",
		last:  "Goot",
		age:   21,
	}
	p3 := person{
		first: "Esya ",
		last:  "Hossain",
		age:   45,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	fmt.Println(p2.first)
	fmt.Println(p3.first)

}
