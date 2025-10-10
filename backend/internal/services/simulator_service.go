package services

import (
	"math/rand"
	"time"

	"github.com/janrusell-dev/brokerx/internal/broker"
)

func StartSimulator(b *broker.Broker, m *MetricsService) {
	topics := []string{"alerts", "orders", "logs"}
	senders := []string{"service-A", "service-B", "service-C"}

	for {
		topic := topics[rand.Intn(len(senders))]
		sender := senders[rand.Intn(len(senders))]

		msg := broker.Message{
			Topic:     topic,
			Sender:    sender,
			Payload:   "Random payload from " + sender,
			Timestamp: time.Now(),
		}

		b.Publish(topic, msg)
		m.RecordMessage(topic, int64(rand.Intn(100)))

		time.Sleep(time.Millisecond * 500)
	}
}
