package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	num, err := adderr(4)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(num)

}

func adderr(n int) (int, error) {
	if n < 2 {
		return 4040, errors.New("THis is not mathch type errors")

	}
	return n, nil
}
