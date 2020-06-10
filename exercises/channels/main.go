package main

import "fmt"

func main() {
	c := make(chan int) // create channel

	go func() { // init goroutine
		c <- 42 // send 42 to channel
	}()

	fmt.Println(<-c) // take the value of channel
}
