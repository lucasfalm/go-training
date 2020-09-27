package main

import "testing"

func TestSaySomethingNow(t *testing.T) {
	expect := "hey"
	result := saySomethingNow()

	if expect != result {
		t.Errorf("Something is wrong here x)")
	}
}

func BenchmarkSaySomethingNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		saySomethingNow()
	}
}

func TestMain(t *testing.T) {}
