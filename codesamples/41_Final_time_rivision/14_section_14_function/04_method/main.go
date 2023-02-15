package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}
type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println(s.first, s.ltk)
}
func main() {
	s1 := secretAgent{
		person: person{
			first: "Nihad",
			last:  "Hossain",
		},
		ltk: true,
	}
	s2 := secretAgent{
		person: person{
			first: "Alex",
			last:  "Goot",
		},
		ltk: false,
	}
	s1.speak()
	s2.speak()

}
