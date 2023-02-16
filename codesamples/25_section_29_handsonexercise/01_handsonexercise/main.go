package main

import (
	"fmt"
	"pkg/mod/github.com/GoesToTwentyOne/HonorableTodd_udemy_Golang_practice/codesamples/25_section_29_handsonexercise/01_handsonexercise/dog"
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
