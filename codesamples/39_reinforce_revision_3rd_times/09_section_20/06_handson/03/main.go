package main

import "fmt"

type human interface {
	speak() string
}
type person struct {
	first string
	last  string
	age   int
}

func (p *person) speak() string {
	return fmt.Sprintf("%s %s %d\n", p.first, p.last, p.age)
}
func saySomething(h human) {
	//fmt.Println(h)
	fmt.Println(h.speak())
}

func main() {
	p := person{
		first: "Nihad",
		last:  "Hossain",
		age:   21,
	}
	saySomething(&p)
	//cannot run
	//	saySomething(p)

}
