package synchronize

import (
	"fmt"
	"sync"
)

func UsingChannels(countTo int) {
	ch := make(chan int, countTo)
	defer close(ch)

	wg := sync.WaitGroup{}

	for number := 0; number <= countTo; number++ {
		wg.Add(1)

		ch <- number

		go printChannel(&wg, ch)
	}

	wg.Wait()
}

func printChannel(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	fmt.Printf("number %v\n", <-ch)
}
