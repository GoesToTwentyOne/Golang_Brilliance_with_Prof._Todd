package main

import (
	"fmt"
)

func main() {
	switch x := "Nihad"; x {
	case "Hagty":
		fmt.Println("I'm Hagty")
	case "Joya":
		fmt.Println("I'm Joya")
	case "Nih":
		fmt.Println("I'm NIh")
	case "Alex", "Dyna", "Melton":
		fmt.Println("We are team HighL")
	case "Nihad", "Jordan", "Sithi":
		fmt.Println("we are team AntiVirus")
	default:
		{
			fmt.Printf("we are doing great,I'm from default")
		}

	}
}
