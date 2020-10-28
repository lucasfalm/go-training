package main

import (
	"fmt"
	"time"
)

func main() {
	one, two, three := myfunc(1, 2, 3)
	fmt.Println(one, two, three)

	switch one {
	case 1002:
		fmt.Println("1002")
		fallthrough
	default:
		fmt.Println("will also pass here")
	}

	c := make(chan bool)

	go usingChannels(c)
	fmt.Println("before goroutine")

	c <- true

	<-c // wait until the channel send one signal back
}

func myfunc(a, b, c int) (one, two, three int) {
	one = a + 1
	changingValues(&one, 1000)

	two = one + 2
	three = *(notafunc())
	return // the returned values were defined explicity
}

func notafunc() (number *int) {
	i := 55
	number = &(i)
	return
}

func changingValues(value *int, newValue int) {
	*value = newValue + *value
	return
}

func usingChannels(c chan bool) {
	if <-c {
		time.Sleep(3 * time.Second)
		fmt.Println("it was true")
	}
	c <- true // send signal back
}
