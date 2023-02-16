package main

import (
	"fmt"
	"pkg/mod/github.com/GoesToTwentyOne/HonorableTodd_udemy_Golang_practice/codesamples/25_section_29_handsonexercise/02_handsonerxercise_02/quote"
	"pkg/mod/github.com/GoesToTwentyOne/HonorableTodd_udemy_Golang_practice/codesamples/25_section_29_handsonexercise/02_handsonerxercise_02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}

}
