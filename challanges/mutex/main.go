package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0

	const gs = 100
	var mu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock() // lock the race
			y := &counter
			runtime.Gosched() // goroutine sleep
			*y++              // increment
			mu.Unlock()       // unlock thread
			wg.Done()         // finish goroutine
		}()
		fmt.Printf("Thread %v and counter is: %v\t\n", i, counter)
	}
	wg.Wait() // wait for all goroutines end

	fmt.Println("The final value of counter is:", counter)
}
