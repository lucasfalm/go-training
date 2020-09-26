package main

import "testing"

func TestSaySomethingNow(t *testing.T) {
	expect := "heya"
	result := saySomethingNow()

	if expect != result {
		t.Errorf("Something is wrong here x)")
	}
}
