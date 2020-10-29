package main

import (
	"fmt"
	"sync"
	"time"
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
		subject = EventSubject{Observers: sync.Map{}}

		observerOne = EventObserver{ID: 1, Time: time.Now()}
		observerTwo = EventObserver{ID: 2, Time: time.Now()}
		event       = Event{Data: DataSchema{Title: "hey you", Body: "wanna have some fun?"}}
	)

	subject.AddListener(&observerOne)
	subject.AddListener(&observerTwo)
	subject.Notify(event)

	<-c // wait until the channel send one signal back
}

type (
	Event struct {
		Data DataSchema
	}

	DataSchema struct {
		Title string
		Body  string
	}

	EventSubject struct {
		Observers sync.Map
	}

	EventObserver struct {
		ID   int
		Time time.Time
	}

	Observer interface {
		NotifyCallback(Event)
	}

	Subject interface {
		AddListener(Observer)
		RemoveListener(Observer)
		Notify(Event)
	}
)

func (eo *EventObserver) NotifyCallback(e Event) {
	fmt.Printf("ok with event %v", e.Data)
}

func (es *EventSubject) AddListener(observer Observer) {
	es.Observers.Store(observer, struct{}{})
}

func (es *EventSubject) RemoveListener(observer Observer) {
	es.Observers.Delete(observer)
}

func (es *EventSubject) Notify(e Event) {
	es.Observers.Range(func(key, value interface{}) (result bool) {
		if key == nil || value == nil {
			result = false
			return
		}

		key.(Observer).NotifyCallback(e)
		result = true
		return
	})
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
