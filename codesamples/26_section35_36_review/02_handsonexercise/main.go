package starting_code

import (
	"fmt"

	"github.com/GoesToTwentyOne/HonorableTodd_udemy_Golang_practice/codesamples/26_section35_36_review/02_handsonexercise/dog"
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
