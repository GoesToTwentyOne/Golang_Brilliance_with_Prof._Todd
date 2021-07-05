package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := `Reinforcedarling345#`
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(password)
	fmt.Println(bs)
	fmt.Println(string(bs))
	// password := `Reinforcedarling345#5`
	err = bcrypt.CompareHashAndPassword(bs, []byte(password))
	if err != nil {
		fmt.Println("You can't log in")
		return
	}
	fmt.Println("you are loged in successful")

}
