package synchronize

import (
	"fmt"
	"sync"
)

func UsingMutex(countTo int) {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	count := 0

	for number := 0; number <= countTo; number++ {
		wg.Add(1)

		go printMutex(&count, &wg, &mu)
	}

	wg.Wait()
}

func printMutex(number *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer mu.Unlock()
	defer wg.Done()

	mu.Lock()

	fmt.Printf("number %v\n", *number)

	*number++
}
