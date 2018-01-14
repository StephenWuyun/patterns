package main

import (
	"fmt"
	"time"
)

//观察者模式
type (
	Event struct {
		Id int64
	}

	//观察者
	Observer interface {
		OnNotify(event Event)
	}

	//被观察者
	Notifier interface {
		Init()
		Register(o Observer)
		UnRegister(o Observer)
		Notify(event Event)
	}

	EventObserver struct {
	}

	EventNotifier struct {
		observers map[Observer]struct{}
	}
)

func (e *EventNotifier) Init() {
	e.observers = make(map[Observer]struct{})
}

func (e *EventNotifier) Register(o Observer) {
	e.observers[o] = struct{}{}
}

func (e *EventNotifier) UnRegister(o Observer) {
	delete(e.observers, o)
}

func (e *EventNotifier) Notify(event Event) {
	for observer, _ := range e.observers {
		observer.OnNotify(event)
	}
}

func (e *EventObserver) OnNotify(event Event) {
	fmt.Printf("receive event, id:%d\n", event.Id)
}

func main() {
	var observer Observer = &EventObserver{}
	var notifier Notifier = &EventNotifier{}
	notifier.Init()
	notifier.Register(observer)
	stop := time.NewTimer(10 * time.Second).C
	tick := time.NewTicker(time.Second).C
	for {
		select {
		case <-stop:
			return
		case t := <-tick:
			notifier.Notify(Event{Id: t.Unix()})
		}
	}
}
