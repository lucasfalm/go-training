package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Starts!")
	wg.Add(1)
	go bar()

	fmt.Println("Waiting")
	wg.Wait()
	fmt.Println("Done!")
}

func bar() {
	for i := 0; i <= 10; i++ {
		fmt.Println("bar?")
	}

	wg.Done()
}
