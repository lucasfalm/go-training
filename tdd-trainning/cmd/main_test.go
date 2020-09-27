package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSaySomethingNow(t *testing.T) {
	var (
		expect      = "hey"
		timeoutChan = make(chan bool)
		result      string
	)

	go func() {
		result = saySomethingNow()
		timeoutChan <- false
	}()

	go func() {
		time.Sleep(5 * time.Second) // if the saySomethingNow took more than 5 seconds, the test will fail (avoid eternal looping)
		timeoutChan <- true
	}()

	if <-timeoutChan {
		t.Error("say something took more than 5 seconds to run")
	}
	assert.True(t, <-timeoutChan)

	assert.NotNil(t, result)
	assert.Equal(t, expect, result)
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
