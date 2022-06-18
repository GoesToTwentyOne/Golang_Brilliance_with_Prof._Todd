package main

import "fmt"

func main() {
	counter := 1
	for i := 65; i <= 99; i++ {
		{
			if i == 65 {
				fmt.Println(i)
				for {
					if counter > 3 {
						break

					}
					fmt.Printf("\t\t%#U\n", i)
					counter++

				}
			}

			if i == 66 {
				counter = 0
				fmt.Println(i)
				for {

					if counter > 3 {
						break

					}
					fmt.Printf("\t\t%#U\n", i)
					counter++

				}
			}
			fmt.Println("---------------------------------------------------------------------------------------------------------------------------")
		}
		{

			if true {
				counter = 0
				fmt.Println(i)
				for {

					if counter > 3 {
						break

					}
					fmt.Printf("\t\t%#U\n", i)
					counter++

				}
			}
		}
	}

}
