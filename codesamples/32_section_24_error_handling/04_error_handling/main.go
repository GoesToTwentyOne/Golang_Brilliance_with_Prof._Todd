package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("F:/GOLANG_PROJECTS_ADVANCED/src/github.com/GoesToTwentyOne/HonorableTodd_udemy_Golang_practice/codesamples/32_section_24_error_handling/03_error_handling/read.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	bs, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
