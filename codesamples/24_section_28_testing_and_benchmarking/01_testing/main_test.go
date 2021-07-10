package main

import "testing"

func TestGoesAdd(t *testing.T) {
	if g := goesAdd(1, 2, 3, 4, 5, 56, 6); g != 21 {
		t.Errorf("i expected this 21 and got this %v", g)
	}

}
