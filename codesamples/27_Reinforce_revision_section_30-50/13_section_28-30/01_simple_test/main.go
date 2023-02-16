package main

import (
	"fmt"
	"pkg/mod/github.com/GoesToTwentyOne/HonorableTodd_udemy_Golang_practice/codesamples/27_Reinforce_revision_section_30-50/13_section_28-30/01_simple_test/simplemath"
)

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	f := simplemath.Add(xi...)
	fmt.Println(f)

}
