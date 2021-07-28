package main

import (
	"fmt"

	"github.com/GoesToTwentyOne/HonorableTodd_udemy_Golang_practice/codesamples/28_again_reinforce_practice_and_crystal_idea_all_topic/12_section_28-29/02_handsonexercise/quote"
	"github.com/GoesToTwentyOne/HonorableTodd_udemy_Golang_practice/codesamples/28_again_reinforce_practice_and_crystal_idea_all_topic/12_section_28-29/02_handsonexercise/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
