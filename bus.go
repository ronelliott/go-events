package events

import (
	"sync"

	"github.com/satori/go.uuid"
)

// An event bus
type Bus interface {
	// Does any necessary cleanup
	Close()

	// Emit the given data on the given channel
	Emit(string, ...interface{})

	// Subscribe the given callback to events emitted on the given channel, returning
	// the uuid for the created subscription. This uuid is needed for removing a subscription
	On(string, Callback) string

	// Remove a subscription with the given id
	Remove(string)
}

// The default event bus implementation
type BusImpl struct {
	// The identifier to channel mapping
	Identifiers map[string]string

	// A lock for modifying the other maps for this bus
	Mutex sync.Mutex

	// The subscription channel mapping callback
	Subscriptions map[string]map[string]Callback

	// The wait group for callbacks for this bus
	WaitGroup sync.WaitGroup
}

// Create a new bus
func NewBus() Bus {
	return &BusImpl{
		Identifiers:   map[string]string{},
		Subscriptions: map[string]map[string]Callback{},
	}
}

// Does any necessary cleanup, waiting for all async callbacks to return before returning
func (bus *BusImpl) Close() {
	bus.WaitAsync()
}

// Emit the given data on the given channel. Callbacks will be called on a
// separate goroutine.
func (bus *BusImpl) Emit(channel string, data ...interface{}) {
	if bus.Subscriptions == nil {
		return
	}

	callbacks, ok := bus.Subscriptions[channel]

	if !ok {
		return
	}

	for _, callback := range callbacks {
		bus.WaitGroup.Add(1)
		go func() {
			defer bus.WaitGroup.Done()
			callback(data...)
		}()
	}
}

// Subscribe the given callback to events emitted on the given channel, returning
// the uuid for the created subscription. This uuid is needed for removing a subscription
func (bus *BusImpl) On(channel string, callback Callback) string {
	bus.Mutex.Lock()
	defer bus.Mutex.Unlock()

	if bus.Identifiers == nil {
		bus.Identifiers = map[string]string{}
	}

	if bus.Subscriptions == nil {
		bus.Subscriptions = map[string]map[string]Callback{}
	}

	if _, ok := bus.Subscriptions[channel]; !ok {
		bus.Subscriptions[channel] = map[string]Callback{}
	}

	var id string

	for {
		id = uuid.NewV4().String()

		if _, ok := bus.Identifiers[id]; !ok {
			break
		}
	}

	bus.Identifiers[id] = channel
	bus.Subscriptions[channel][id] = callback
	return id
}

// Remove a subscription with the given id
func (bus *BusImpl) Remove(id string) {
	if bus.Identifiers == nil {
		return
	}

	if bus.Subscriptions == nil {
		return
	}

	bus.Mutex.Lock()
	defer bus.Mutex.Unlock()

	channel := bus.Identifiers[id]
	callbacks, ok := bus.Subscriptions[channel]

	if !ok {
		return
	}

	delete(bus.Identifiers, id)
	delete(callbacks, id)

	if len(callbacks) == 0 {
		delete(bus.Subscriptions, channel)
	}
}

// Waits for existing callbacks to finish, then returns
func (bus *BusImpl) WaitAsync() {
	bus.WaitGroup.Wait()
}
