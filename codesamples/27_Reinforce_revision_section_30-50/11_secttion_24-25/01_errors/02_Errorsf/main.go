package main

import (
	"fmt"
	"log"
)

func main() {
	num, err := adderr(1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(num)

}

func adderr(n int) (int, error) {
	if n < 2 {
		return 4040, fmt.Errorf("%d n is not sufficiant for the program", n)

	}
	return n, nil
}
