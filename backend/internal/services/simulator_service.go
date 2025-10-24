package services

import (
	"math/rand"
	"time"

	"github.com/janrusell-dev/brokerx/internal/broker"
	"github.com/janrusell-dev/brokerx/internal/utils"
)

func StartSimulator(b *broker.Broker, m *MetricsService) {
	topics := []string{"alerts", "orders", "logs"}
	senders := []string{"service-A", "service-B", "service-C"}

	// Seed random number generator
	utils.LogInfo("Message simulator started - publishing every 500ms")

	for {
		/// Pick random topic and sender
		topic := topics[rand.Intn(len(topics))]
		sender := senders[rand.Intn(len(senders))]

		msg := broker.Message{
			Topic:     topic,
			Sender:    sender,
			Payload:   utils.GeneratePayload(topic, sender),
			Timestamp: time.Now(),
		}

		// Measure publish latency
		start := time.Now()
		b.Publish(topic, msg)
		latency := time.Since(start).Microseconds()
		// Record metrics
		m.RecordMessage(topic, latency)

		time.Sleep(time.Millisecond * 500)
	}
}
