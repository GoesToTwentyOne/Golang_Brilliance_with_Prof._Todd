package main

import (
	"fmt"

	"github.com/GoesToTwentyOne/HonorableTodd_udemy_Golang_practice/codesamples/27_Reinforce_revision_section_30-50/13_section_28-30/07_reinforce_handsonexercise/02_hansdsonexercise/quote"
	"github.com/GoesToTwentyOne/HonorableTodd_udemy_Golang_practice/codesamples/27_Reinforce_revision_section_30-50/13_section_28-30/07_reinforce_handsonexercise/02_hansdsonexercise/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
