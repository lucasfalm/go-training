package main

import (
	"testing"
	"time"
)

func TestSaySomethingNow(t *testing.T) {
	expect := "hey"
	timeoutChan := make(chan bool)
	var result string

	go func() {
		result = saySomethingNow()
		timeoutChan <- false
	}()

	go func() {
		time.Sleep(5 * time.Second)
		timeoutChan <- true
	}()

	if <-timeoutChan {
		t.Error("say something took more than 5 seconds to run")
	}

	if expect != result {
		t.Errorf("Something is wrong here x)")
	}
}

/*
We use benchmark to validate which solution is more
performatic
*/

func BenchmarkSaySomethingNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		saySomethingNow()
	}
}

func TestMain(t *testing.T) {}
