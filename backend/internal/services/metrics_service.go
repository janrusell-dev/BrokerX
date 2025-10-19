package services

import (
	"sync"
	"time"
)

type Metrics struct {
	TotalMessages     int64                    `json:"totalMessages"`
	AvgLatency        float64                  `json:"avgLatency"`
	MessagePerTopic   map[string]int           `json:"topicMetrics"`
	ActiveSubscribers int64                    `json:"activeSubscribers"`
	LatencyHistory    []map[string]interface{} `json:"latencyHistory"`
}

type MetricsService struct {
	mu   sync.Mutex
	data Metrics
}

func NewMetricsService() *MetricsService {
	return &MetricsService{
		data: Metrics{
			MessagePerTopic: make(map[string]int),
		},
	}
}

func (m *MetricsService) RecordMessage(topic string, latencyMs int64) {
	m.mu.Lock()
	defer m.mu.Unlock()

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
