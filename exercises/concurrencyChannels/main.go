package main

import "fmt"

func main() {
	c := make(chan int, 2)
	// send
	go foo(c) // init goroutine
	// receive
	bar(c) // not need goroutine because its like wg.wait() for block channel

	fmt.Println("Exiting...")
}

func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i // pass i value to channel, 100 times
	}

	close(c) // close the channel
}

func bar(c <-chan int) {
	for v := range c { // range over the channel
		fmt.Println(v)
	}
}
