package main

import "fmt"

func main() {
	c := make(chan int)
	// send
	go foo(c) // init goroutine
	// receive
	bar(c) // not need goroutine because its like wg.wait() for block channel

	fmt.Println("Exiting...")
}

func foo(c chan<- int) {
	c <- 55
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}
