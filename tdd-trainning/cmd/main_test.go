package main

import "testing"

func TestSaySomethingNow(t *testing.T) {
	expect := "hey"
	result := saySomethingNow()

	if expect != result {
		t.Errorf("Something is wrong here x)")
	}
}

func TestMain(t *testing.T) {}
