package main

import "fmt"

type person struct {
	name    string
	age     int
	address string
}

func (p *person) speak() string {
	return fmt.Sprintf("Name :%s Address: %s", p.name, p.address)
}

type human interface {
	speak() string
}

func saying(h human) {
	fmt.Println(h.speak())
}

func main() {
	p1 := person{
		name:    "Nihad Hossain",
		age:     21,
		address: "Hazipara Thakurgaon",
	}

	p2 := person{
		name:    "Seniya Aliya",
		age:     22,
		address: "Canada",
	}

	p3 := person{
		name:    "Alex Goot",
		age:     21,
		address: "USA",
	}

	// fmt.Println("Print all method set ")
	fmt.Println(p1.speak())
	fmt.Println(p2.speak())
	fmt.Println(p3.speak())

	// fmt.Println("Done")
	// saying(&p1)
	// saying(&p2)
	// saying(&p3)

}
