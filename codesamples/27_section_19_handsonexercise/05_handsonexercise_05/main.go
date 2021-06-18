package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func (p user) String() string {
	return fmt.Sprintf("%s %s %d %s", p.First, p.Last, p.Age, p.Sayings)
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []user

func (bl ByLast) Len() int           { return len(bl) }
func (bl ByLast) Swap(i, j int)      { bl[i], bl[j] = bl[j], bl[i] }
func (bl ByLast) Less(i, j int) bool { return bl[i].Last < bl[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}
	for _, usersvalue := range users {
		fmt.Println(usersvalue.First, "\t \t", usersvalue.Last, "\t \t", usersvalue.Age)
		fmt.Println("---------------------------------------------------------------------------------------")
		for _, sayvalue := range usersvalue.Sayings {
			fmt.Println(sayvalue)

		}
		fmt.Printf("\n \n\n")
	}

	// your code goes here
	fmt.Println("--------------------------------after by age sort -----------------------------------------------")
	sort.Sort(ByAge(users))
	for _, usersvalue := range users {
		fmt.Println(usersvalue.First, "\t \t", usersvalue.Last, "\t \t", usersvalue.Age)
		fmt.Println("---------------------------------------------------------------------------------------")
		for _, sayvalue := range usersvalue.Sayings {
			fmt.Println(sayvalue)

		}
		fmt.Printf("\n \n\n")
	}

	fmt.Println("--------------------------------after by Lastname sort -----------------------------------------------")
	sort.Sort(ByLast(users))
	for _, usersvalue := range users {
		fmt.Println(usersvalue.First, "\t \t", usersvalue.Last, "\t \t", usersvalue.Age)
		fmt.Println("---------------------------------------------------------------------------------------")
		for _, sayvalue := range usersvalue.Sayings {
			fmt.Println(sayvalue)

		}
		fmt.Printf("\n \n\n")
	}

}
