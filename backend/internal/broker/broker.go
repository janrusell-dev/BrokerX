// internal/broker/broker.go
package broker

import (
	"sync"
	"time"
)

type Message struct {
	Topic     string                 `json:"topic"`
	Sender    string                 `json:"sender"`
	Payload   map[string]interface{} `json:"payload"`
	Timestamp time.Time              `json:"timestamp"`
}

type Broker struct {
	mu         sync.RWMutex
	topics     map[string][]chan Message
	topicStats map[string]*TopicStats
}

type TopicStats struct {
	MessageCount  int64
	LastPublished time.Time
}

func NewBroker() *Broker {
	return &Broker{
		topics:     make(map[string][]chan Message),
		topicStats: make(map[string]*TopicStats),
	}
}

// Publish sends a message to all subscribers of a topic
func (b *Broker) Publish(topic string, msg Message) {
	b.mu.RLock()
	subscribers, exists := b.topics[topic]
	b.mu.RUnlock()

	if !exists || len(subscribers) == 0 {
		return
	}

	// Update stats
	b.mu.Lock()
	if b.topicStats[topic] == nil {
		b.topicStats[topic] = &TopicStats{}
	}
	b.topicStats[topic].MessageCount++
	b.topicStats[topic].LastPublished = time.Now()
	b.mu.Unlock()

	// Broadcast to all subscribers (non-blocking)
	for _, ch := range subscribers {
		select {
		case ch <- msg:
		default:
			// Skip slow subscribers to prevent blocking
		}
	}
}

// Subscribe creates a new subscriber channel for a topic
func (b *Broker) Subscribe(topic string) chan Message {
	b.mu.Lock()
	defer b.mu.Unlock()

	ch := make(chan Message, 100) // Buffered channel

	if b.topics[topic] == nil {
		b.topics[topic] = []chan Message{}
	}
	b.topics[topic] = append(b.topics[topic], ch)

	if b.topicStats[topic] == nil {
		b.topicStats[topic] = &TopicStats{}
	}

	return ch
}

// Unsubscribe removes a subscriber channel from a topic
func (b *Broker) Unsubscribe(topic string, ch chan Message) {
	b.mu.Lock()
	defer b.mu.Unlock()

	subscribers := b.topics[topic]
	for i, subscriber := range subscribers {
		if subscriber == ch {
			// Remove this subscriber
			b.topics[topic] = append(subscribers[:i], subscribers[i+1:]...)
			close(ch)
			break
		}
	}

	// Clean up empty topics
	if len(b.topics[topic]) == 0 {
		delete(b.topics, topic)
	}
}

// GetTopics returns all active topics
func (b *Broker) GetTopics() []string {
	b.mu.RLock()
	defer b.mu.RUnlock()

	topics := make([]string, 0, len(b.topics))
	for topic := range b.topics {
		topics = append(topics, topic)
	}
	return topics
}

// GetTopicInfo returns detailed info about a topic
func (b *Broker) GetTopicInfo(topic string) map[string]interface{} {
	b.mu.RLock()
	defer b.mu.RUnlock()

	info := map[string]interface{}{
		"exists":        false,
		"subscribers":   0,
		"messageCount":  int64(0),
		"lastPublished": nil,
	}

	if subscribers, exists := b.topics[topic]; exists {
		info["exists"] = true
		info["subscribers"] = len(subscribers)

		if stats, ok := b.topicStats[topic]; ok {
			info["messageCount"] = stats.MessageCount
			if !stats.LastPublished.IsZero() {
				info["lastPublished"] = stats.LastPublished
			}
		}
	}

	return info
}

// GetAllTopicsInfo returns info for all topics
func (b *Broker) GetAllTopicsInfo() []map[string]interface{} {
	b.mu.RLock()
	defer b.mu.RUnlock()

	result := make([]map[string]interface{}, 0, len(b.topics))

	for topic := range b.topics {
		info := map[string]interface{}{
			"topic":         topic,
			"subscribers":   len(b.topics[topic]),
			"messageCount":  int64(0),
			"lastPublished": nil,
		}

		if stats, ok := b.topicStats[topic]; ok {
			info["messageCount"] = stats.MessageCount
			if !stats.LastPublished.IsZero() {
				info["lastPublished"] = stats.LastPublished
			}
		}

		result = append(result, info)
	}

	return result
}
