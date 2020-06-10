package main

import "fmt"

func main() {
	c := make(chan int, 1) // create channel with 1 goroutine
	c <- 42
	//go func() { // init goroutine
	//	c <- 42 // send 42 to channel
	//}()

	fmt.Println(<-c) // take the value of channel
}
