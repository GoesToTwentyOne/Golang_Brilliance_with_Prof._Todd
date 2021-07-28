package main

import (
	"fmt"

	"github.com/GoesToTwentyOne/HonorableTodd_udemy_Golang_practice/codesamples/28_again_reinforce_practice_and_crystal_idea_all_topic/12_section_28-29/03_handsonexercise/mymath"
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
