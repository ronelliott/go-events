package events

import "sync"

// An event bus
type Bus interface {
	// Does any necessary cleanup
	Close()

	// Subscribe the given callback to all events emitted on all channels, returning
	// the uuid for the created subscription. This uuid is needed for removing a subscription.
	// The callback will be called after any subscribed callbacks
	After(MiddlewareCallback)

	// Subscribe the given callback to all events emitted on all channels, returning
	// the uuid for the created subscription. This uuid is needed for removing a subscription.
	// The callback will be called before any subscribed callbacks
	Before(MiddlewareCallback)

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
	// The "after" message callbacks
	Afters map[string]MiddlewareCallback

	// The "before" message callbacks
	Befores map[string]MiddlewareCallback

	// The identifier to channel mapping
	Identifiers map[string]string

	// A lock for modifying the other maps for this bus
	Mutex sync.Mutex

	// The subscription channel mapping callback
	Subscriptions map[string]map[string]Callback
}

// Create a new bus
func NewBus() Bus {
	return &BusImpl{
		Afters:        map[string]MiddlewareCallback{},
		Befores:       map[string]MiddlewareCallback{},
		Identifiers:   map[string]string{},
		Subscriptions: map[string]map[string]Callback{},
	}
}

// Does any necessary cleanup
func (bus *BusImpl) Close() {}

// Subscribe the given callback to all events emitted on all channels, returning
// the uuid for the created subscription. This uuid is needed for removing a subscription.
// The callback will be called after any subscribed callbacks
func (bus *BusImpl) After(callback MiddlewareCallback) {
	bus.Mutex.Lock()
	defer bus.Mutex.Unlock()

	if bus.Afters == nil {
		bus.Afters = map[string]MiddlewareCallback{}
	}

	var id string

	for {
		id = uuid()

		if _, ok := bus.Afters[id]; !ok {
			break
		}
	}

	bus.Afters[id] = callback
}

// Subscribe the given callback to all events emitted on all channels, returning
// the uuid for the created subscription. This uuid is needed for removing a subscription.
// The callback will be called before any subscribed callbacks
func (bus *BusImpl) Before(callback MiddlewareCallback) {
	bus.Mutex.Lock()
	defer bus.Mutex.Unlock()

	if bus.Befores == nil {
		bus.Befores = map[string]MiddlewareCallback{}
	}

	var id string

	for {
		id = uuid()

		if _, ok := bus.Befores[id]; !ok {
			break
		}
	}

	bus.Befores[id] = callback
}

// Emit the given data on the given channel. Callbacks will be called on a
// separate goroutine.
func (bus *BusImpl) Emit(channel string, data ...interface{}) {
	if bus.Subscriptions == nil {
		return
	}

	callbacks, ok := bus.Subscriptions[channel]

	wg := sync.WaitGroup{}

	if bus.Befores != nil {
		for _, callback := range bus.Befores {
			wg.Add(1)
			go func() {
				defer wg.Done()
				callback(channel, data...)
			}()
		}

		wg.Wait()
		wg = sync.WaitGroup{}
	}

	if ok {
		for _, callback := range callbacks {
			wg.Add(1)
			go func() {
				defer wg.Done()
				callback(data...)
			}()
		}

		wg.Wait()
	}

	if bus.Afters != nil {
		wg = sync.WaitGroup{}

		for _, callback := range bus.Afters {
			wg.Add(1)
			go func() {
				defer wg.Done()
				callback(channel, data...)
			}()
		}

		wg.Wait()
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
		id = uuid()

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
