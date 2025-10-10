package broker

import "sync"

type Broker struct {
	mu          sync.RWMutex
	subscribers map[string][]chan Message
}

func NewBroker() *Broker {
	return &Broker{
		subscribers: make(map[string][]chan Message),
	}
}

func (b *Broker) Publish(topic string, msg Message) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if subs, ok := b.subscribers[topic]; ok {
		for _, ch := range subs {
			go func(c chan Message) {
				c <- msg
			}(ch)
		}
	}
}

func (b *Broker) Subscribe(topic string) <-chan Message {
	b.mu.Lock()
	defer b.mu.Unlock()

	ch := make(chan Message, 100)
	b.subscribers[topic] = append(b.subscribers[topic], ch)
	return ch
}
