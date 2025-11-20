package services

import (
	"sync"
	"time"

	"github.com/brokerx/internal/broker"
)

type Metrics struct {
	TotalMessages     int64                    `json:"totalMessages"`
	AvgLatency        float64                  `json:"avgLatency"`
	MessagePerTopic   map[string]int           `json:"topicMetrics"`
	ActiveSubscribers int64                    `json:"activeSubscribers"`
	LatencyHistory    []map[string]interface{} `json:"latencyHistory"`
	LastReset         time.Time                `json:"lastReset"`
}

type MetricsService struct {
	mu           sync.Mutex
	data         Metrics
	startTime    time.Time
	maxMessages  int
	messageQueue []broker.Message
}

func NewMetricsService(maxMessages int) *MetricsService {
	return &MetricsService{
		data: Metrics{
			MessagePerTopic: make(map[string]int),
			LastReset:       time.Now(),
		},
		startTime:   time.Now(),
		maxMessages: maxMessages,
	}
}

func (m *MetricsService) RecordMessage(topic string, latencyMs int64, msg broker.Message) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.messageQueue = append(m.messageQueue, msg)
	if len(m.messageQueue) > m.maxMessages {
		oldMsg := m.messageQueue[0]
		m.messageQueue = m.messageQueue[1:]

		if m.data.MessagePerTopic[oldMsg.Topic] > 0 {
			m.data.MessagePerTopic[oldMsg.Topic]--
		}
		m.data.TotalMessages--
	}

	m.data.TotalMessages++
	m.data.MessagePerTopic[topic]++
	m.data.AvgLatency = (m.data.AvgLatency + float64(latencyMs)) / 2

	// Keep up to the last 20 latency samples
	if len(m.data.LatencyHistory) >= 20 {
		m.data.LatencyHistory = m.data.LatencyHistory[1:]
	}

	m.data.LatencyHistory = append(m.data.LatencyHistory, map[string]interface{}{
		"timestamp": time.Now().Format(time.RFC3339),
		"latency":   latencyMs,
	})
}

func (m *MetricsService) GetMetrics() Metrics {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.data

}
func (m *MetricsService) IncrementSubscribers() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data.ActiveSubscribers++
}

func (m *MetricsService) DecrementSubscribers() {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data.ActiveSubscribers > 0 {
		m.data.ActiveSubscribers--
	}
}

func (m *MetricsService) ResetMetrics() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.data.TotalMessages = 0
	m.data.AvgLatency = 0
	m.data.MessagePerTopic = make(map[string]int)
	m.data.ActiveSubscribers = 0
	m.data.LatencyHistory = []map[string]interface{}{}
	m.data.LastReset = time.Now()
}

func (m *MetricsService) MessageRate() float64 {
	m.mu.Lock()
	defer m.mu.Unlock()

	duration := time.Since(m.data.LastReset).Seconds()
	if duration == 0 {
		return 0
	}
	return float64(m.data.TotalMessages) / duration
}
