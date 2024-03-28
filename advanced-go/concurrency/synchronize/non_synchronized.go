package synchronize

import (
	"fmt"
	"sync"
)

func NonSynchronized(countTo int) {
	wg := &sync.WaitGroup{}

	for number := 0; number <= countTo; number++ {
		wg.Add(1)

		go printNonSynchronized(number, wg)
	}

	wg.Wait()
}

func printNonSynchronized(number int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("number %v\n", number)
}
