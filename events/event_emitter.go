package events

import "sync"

type handler = func(...interface{})

// event emitter
type EventEmitter struct {
	sync.Mutex
	handlers map[string][]handler // 事件处理队列
}

// create a new event emitter
func NewEventEmitter() *EventEmitter {
	return &EventEmitter{
		handlers: make(map[string][]handler),
	}
}

// add event handler
func (p *EventEmitter) On(event string, handler handler) {
	p.Lock()
	p.handlers[event] = append(p.handlers[event], handler)
	p.Unlock()
}

// remove event handler
func (p *EventEmitter) Off(event string) {
	p.Lock()
	delete(p.handlers, event)
	p.Unlock()
}

// trigger event
func (p EventEmitter) Emit(event string, args ...interface{}) {
	for i := range p.handlers[event] {
		go p.handlers[event][i](args...)
	}
}
