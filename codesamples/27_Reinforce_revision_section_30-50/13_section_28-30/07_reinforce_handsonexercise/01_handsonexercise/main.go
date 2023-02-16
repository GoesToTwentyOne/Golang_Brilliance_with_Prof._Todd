package main

import (
	"fmt"
	"pkg/mod/github.com/GoesToTwentyOne/HonorableTodd_udemy_Golang_practice/codesamples/27_Reinforce_revision_section_30-50/13_section_28-30/07_reinforce_handsonexercise/01_handsonexercise/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
