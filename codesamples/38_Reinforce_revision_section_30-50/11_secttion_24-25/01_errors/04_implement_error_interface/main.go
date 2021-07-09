package main

import "fmt"

type cuserr struct {
	math    string
	science string
	err     error
}

func (cs cuserr) Error() string {
	return fmt.Sprintf("subject error %s science error %s custom err %s", cs.math, cs.science, cs.err)
}

func main() {
	num, err := numcustomerr(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)

}
func numcustomerr(n int) (int, error) {
	if n < 2 {
		myerr := fmt.Errorf("N is not sufficiant for the program n=%d", n)
		return 0, cuserr{"sin--------", "science ---------------", myerr}
	}
	return n, nil

}
