package main

import (
	"fmt"
	mymath "pkg/mod/github.com/GoesToTwentyOne/HonorableTodd_udemy_Golang_practice/codesamples/27_Reinforce_revision_section_30-50/13_section_28-30/07_reinforce_handsonexercise/03_handsonexercise/Mymath"
)

func main() {
	xxi := gen()
	for _, v := range xxi {
		fmt.Println(mymath.CenteredAvg(v))
	}
}

func gen() [][]int {
	a := []int{1, 4, 6, 8, 100}
	b := []int{0, 8, 10, 1000}
	c := []int{9000, 4, 10, 8, 6, 12}
	d := []int{123, 744, 140, 200}
	e := [][]int{a, b, c, d}
	return e
}
