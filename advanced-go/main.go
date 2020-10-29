package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/flucas97/go-trainning/advanced-go/observer"
)

/*
observer

Subject - sync.map{} - add,remove listener, notify
Observer - notifyCallback


*/

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

	var (
		subject     = observer.EventSubject{Observers: sync.Map{}}
		observerOne = observer.EventObserver{ID: 1, Time: time.Now()}
		observerTwo = observer.EventObserver{ID: 2, Time: time.Now()}
		event       = observer.Event{Data: observer.DataSchema{Title: "hey you", Body: "wanna have some fun?"}}
	)

	subject.AddListener(&observerOne)
	subject.AddListener(&observerTwo)
	subject.Notify(event)

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
	number = &i
	return
}

func changingValues(value *int, newValue int) {
	*value = newValue + *value
	return
}

func usingChannels(c chan bool) {
	if <-c { // it get the value from the channel
		time.Sleep(3 * time.Second)
		fmt.Println("it was true")
	}
	c <- true // send signal back
}
