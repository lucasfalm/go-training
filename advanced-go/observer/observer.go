package observer

import (
	"fmt"
	"sync"
	"time"
)

/*
Observer Design Pattern
------------------------
Subject - sync.map{} - add,remove listener, notify
Observer - notifyCallback
*/

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
	fmt.Printf("[Event Observer ID: %v] receive title:'%v' and body: %v \n", eo.ID, e.Data.Title, e.Data.Body)
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
