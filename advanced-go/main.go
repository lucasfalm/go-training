package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/flucas97/go-trainning/advanced-go/observer"
)

var (
	MyFuncSubject = observer.EventSubject{Observers: sync.Map{}}
)

func main() {
	var (
		observerOne = observer.EventObserver{ID: 1, Time: time.Now()}
		observerTwo = observer.EventObserver{ID: 2, Time: time.Now()}
	)

	MyFuncSubject.AddListener(&observerOne)
	MyFuncSubject.AddListener(&observerTwo)

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
	event := observer.Event{Data: observer.DataSchema{Title: "hey you", Body: "wanna have some fun?"}}
	MyFuncSubject.Notify(event)

	one = a + 1
	changingValues(&one, 1000)
	two = one + 2
	three = *(notafunc())
	return // the returned values were defined explicity
}

func notafunc() (number *int) {
	i := rand.Intn(100*int(time.Millisecond)) / 100000
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
