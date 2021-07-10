package main

import (
	"fmt"

	"github.com/GoesToTwentyOne/HonorableTodd_udemy_Golang_practice/codesamples/26_section35_36_review/03_handsonerxercise_02/quote"
	"github.com/GoesToTwentyOne/HonorableTodd_udemy_Golang_practice/codesamples/26_section35_36_review/03_handsonerxercise_02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
