package main

import "fmt"

type person struct {
	first string
	last  string
}
type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println(s.first, s.last, "secret agent")
}
func (p person) speak() {
	fmt.Println(p.first, p.last, "person")
}

type human interface {
	speak()
}

func bar(h human) {

	switch h.(type) {
	case person:
		fmt.Println("person bar  type ", h.(person).first)
	case secretAgent:
		fmt.Println("secretagent bar  type ", h.(secretAgent).first)
	}
	fmt.Println(h)

}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "Nihad",
			last:  "Hossain",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			first: "Henry ",
			last:  "cotiya",
		},
		ltk: true,
	}
	p := person{
		first: "Navida",
		last:  "white",
	}
	bar(sa1)
	bar(sa2)
	bar(p)

}
