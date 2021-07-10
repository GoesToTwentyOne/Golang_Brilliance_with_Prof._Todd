package main

import (
	"fmt"
	"strings"

	"github.com/GoesToTwentyOne/HonorableTodd_udemy_Golang_practice/codesamples/37_section35_36_review/01_benchmarking_testing_and_compare_join_func_from_string_package/cat"
)

const s = "My name is Nihad Hossain "

func main() {
	xs := strings.Split(s, " ")
	for _, value := range xs {
		fmt.Println(value)

	}
	fmt.Printf("%s \n \n", cat.Mystring(xs))
	fmt.Printf("%s \n", cat.Join(xs))
	fmt.Println(len(s))
	fmt.Println(len(cat.Mystring(xs)))
	fmt.Println(len(cat.Join(xs)))

}
