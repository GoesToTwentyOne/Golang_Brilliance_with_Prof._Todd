package vero

import (
	"fmt"
	"strings"
	"testing"
)

const s = "We ask ourselves, Who am I to be brilliant, gorgeous, talented, fabulous? Actually, who are you not to be? Your playing small does not serve the world. There is nothing enlightened about shrinking so that other people won't feel insecure around you. We are all meant to shine, as children do. We were born to make manifest the glory that is within us. It's not just in some of us; it's in everyone. And as we let our own light shine, we unconsciously give other people permission to do the same. As we are liberated from our own fear, our presence automatically liberates others. - Marianne Williamson"

var xs = strings.Split(s, " ")

func TestConcat(t *testing.T) {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	s = Concat(xs)
	if s != "Shaken not stirred" {
		t.Error("got", s, "want", "Shaken not stirred")
	}
}

func TestJoin(t *testing.T) {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	s = Join(xs)
	if s != "Shaken not stirred" {
		t.Error("got", s, "want", "Shaken not stirred")
	}
}

func ExampleConcat() {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	fmt.Println(Concat(xs))
	// Output:
	// Shaken not stirred
}

func ExampleJoin() {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	fmt.Println(Join(xs))
	// Output:
	// Shaken not stirred
}

//----------------------------------------------------------------------------------------
func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Concat(xs)
	}

}
func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Join(xs)
	}

}
