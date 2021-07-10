package main

import "fmt"

type person struct {
	first string
	last  string
}
type ceo struct {
	person
	ceo_qualufication string
}

func (c ceo) treat() {
	fmt.Println(c.first)

}

func main() {
	c1 := ceo{
		person: person{
			first: "Nihad",
			last:  "Hossain",
		},
		ceo_qualufication: "PHD",
	}
	c2 := ceo{
		person: person{
			first: "Thani",
			last:  "Hossain",
		},
		ceo_qualufication: "PHD",
	}
	c1.treat()
	c2.treat()

}
