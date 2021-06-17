package main

import (
	"fmt"
	"sort"
)

type Book struct {
	Name    string
	Page    int
	Writter string
	Serial  int
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type Byserial []Book

func (a Byserial) Len() int           { return len(a) }
func (a Byserial) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Byserial) Less(i, j int) bool { return a[i].Serial < a[j].Serial }

func (b Book) String() string {
	return fmt.Sprintf("%s: %d %s %d", b.Name, b.Page, b.Writter, b.Serial)
}

func main() {
	b1 := Book{
		Name:    "Helllo Hii",
		Page:    500,
		Writter: "Angela --",
		Serial:  5}
	b2 := Book{
		Name:    "Helllo Hii 2",
		Page:    500,
		Writter: "Angela --2",
		Serial:  255,
	}
	b3 := Book{
		Name:    "Helllo Hii--3",
		Page:    500,
		Writter: "Angela --3",
		Serial:  55,
	}
	b4 := Book{
		Name:    "Helllo Hii -- 4",
		Page:    500,
		Writter: "Angela --  4",
		Serial:  785,
	}
	b5 := Book{
		Name:    "Helllo Hii -- 5",
		Page:    500,
		Writter: "Angela -- 5",
		Serial:  455,
	}
	all := []Book{b1, b2, b3, b4, b5}
	sort.Sort(Byserial(all))
	fmt.Println(all)

}
