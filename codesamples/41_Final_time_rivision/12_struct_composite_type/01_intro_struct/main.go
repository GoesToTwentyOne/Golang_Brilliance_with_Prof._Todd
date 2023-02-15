package main

import (
	"fmt"
)

type person struct {
	first_name string
	last_name  string
	skills     string
	core_value int
	age        int
}

func main() {
	p := person{
		first_name: "Nihad",
		last_name:  "Hossain",
		skills:     "Golang,C++,C",
		core_value: 01,
		age:        21,
	}
	p1 := person{
		first_name: "Yerti",
		last_name:  "Khan",
		skills:     "Golang,C++,C",
		core_value: 01,
		age:        21,
	}
	fmt.Println(p)
	fmt.Println(p1)
	fmt.Println(p.first_name, p.skills)
	fmt.Println(p1.first_name, p1.skills)

}
