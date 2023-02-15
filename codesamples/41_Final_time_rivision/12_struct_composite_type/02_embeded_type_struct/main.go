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
type secretAgent struct {
	person
	ltk                bool
	survaival_training bool
}

func main() {
	s1 := secretAgent{
		person: person{
			first_name: "Nihad",
			last_name:  "Hossain",
			skills:     "Golang,C++,C,Survival Experise",
			age:        22,
		},
		ltk:                true,
		survaival_training: true,
	}
	fmt.Println(s1)

}
