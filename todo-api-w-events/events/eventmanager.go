package events

import (
	"context"
	"fmt"
	"log"
)

type ToDoEventManager struct {
	ctx      context.Context
	cancel   context.CancelFunc
	queue    chan *ToDoEvent
	isActive bool
}

func NewToDoEventManager() *ToDoEventManager {
	return &ToDoEventManager{
		ctx:      nil,
		cancel:   nil,
		queue:    make(chan *ToDoEvent),
		isActive: false,
	}
}

func (em *ToDoEventManager) Start() {
	if !em.isActive {
		em.ctx, em.cancel = context.WithCancel(context.Background())
		em.isActive = true
		go em.eventLoop()
	}
}

func (em *ToDoEventManager) eventLoop() {
	log.Println("Starting Event Loop...")
	for {
		select {
		case <-em.ctx.Done():
			log.Println("Stopping Event Manager...")
			return
		case event := <-em.queue:
			log.Printf("\n--> Received Event: %+v\n", event.EventData)
			em.processEvent(event)
		}
	}
}

func (em *ToDoEventManager) Stop() {
	if em.isActive {
		em.cancel()
		em.isActive = false
	}
}

func (em *ToDoEventManager) Notify(event *ToDoEvent) {
	if em.isActive {
		em.queue <- event
	}
}

func (em *ToDoEventManager) processEvent(event *ToDoEvent) {
	switch event.EventID {
	case ToDoQueryEvent:
		em.processQueryEvent(event)
	case ToDoAddEvent:
		em.processAddEvent(event)
	case ToDoUpdateEvent:
		em.processUpdateEvent(event)
	case ToDoDeleteEvent:
		em.processDeleteEvent(event)
	case ToDoErrorEvent:
		em.processErrorEvent(event)
	}
}

func (em *ToDoEventManager) processQueryEvent(event *ToDoEvent) {
	fmt.Println("Processing Query Event")
}

func (em *ToDoEventManager) processAddEvent(event *ToDoEvent) {
	fmt.Println("Processing Add Event")
}

func (em *ToDoEventManager) processUpdateEvent(event *ToDoEvent) {
	fmt.Println("Processing Update Event")
}

func (em *ToDoEventManager) processDeleteEvent(event *ToDoEvent) {
	fmt.Println("Processing Delete Event")
}

func (em *ToDoEventManager) processErrorEvent(event *ToDoEvent) {
	fmt.Println("Processing Error Event")
}
