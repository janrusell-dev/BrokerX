package services

import (
	"math/rand"
	"time"

	"github.com/brokerx/internal/broker"
	"github.com/brokerx/internal/utils"
)

type Simulator struct {
	broker  *broker.Broker
	metrics *MetricsService
	ticker  *time.Ticker
	quit    chan struct{}
}

func NewSimulator(b *broker.Broker, m *MetricsService) *Simulator {
	return &Simulator{
		broker:  b,
		metrics: m,
		quit:    make(chan struct{}),
	}
}

func (s *Simulator) Start() {
	s.ticker = time.NewTicker(500 * time.Millisecond)

	go func() {

		topics := []string{"alerts", "orders", "logs"}
		senders := []string{"service-A", "service-B", "service-C"}

		utils.LogInfo("Message simulator started - publishing every 500ms")
		for {
			select {
			case <-s.quit:
				s.ticker.Stop()
				return
			case <-s.ticker.C:
				topic := topics[rand.Intn(len(topics))]
				sender := senders[rand.Intn(len(senders))]
				msg := broker.Message{
					Topic:     topic,
					Sender:    sender,
					Payload:   utils.GeneratePayload(topic, sender),
					Timestamp: time.Now(),
				}

				start := time.Now()
				s.broker.Publish(topic, msg)
				latency := time.Since(start).Microseconds()
				s.metrics.RecordMessage(topic, latency, msg)
			}
		}
	}()
}

func (s *Simulator) Stop() {
	close(s.quit)
}

// // for local testing purposes only
// func StartSimulatorLocal(b *broker.Broker, m *MetricsService) {
// 	topics := []string{"alerts", "orders", "logs"}
// 	senders := []string{"service-A", "service-B", "service-C"}

// 	// Seed random number generator
// 	utils.LogInfo("Message simulator started - publishing every 500ms")

// 	for {
// 		/// Pick random topic and sender
// 		topic := topics[rand.Intn(len(topics))]
// 		sender := senders[rand.Intn(len(senders))]

// 		msg := broker.Message{
// 			Topic:     topic,
// 			Sender:    sender,
// 			Payload:   utils.GeneratePayload(topic, sender),
// 			Timestamp: time.Now(),
// 		}

// 		// Measure publish latency
// 		start := time.Now()
// 		b.Publish(topic, msg)
// 		latency := time.Since(start).Microseconds()
// 		// Record metrics
// 		m.RecordMessage(topic, latency, msg)

// 		time.Sleep(time.Millisecond * 500)
// 	}
// }

// func StartSimulator(b *broker.Broker, m *MetricsService) {
// 	topics := []string{"alerts", "orders", "logs"}
// 	senders := []string{"service-A", "service-B", "service-C"}

// 	// Seed random number generator
// 	utils.LogInfo("Message simulator started - publishing every 500ms")

// 	for {
// 		/// Pick random topic and sender
// 		topic := topics[rand.Intn(len(topics))]
// 		sender := senders[rand.Intn(len(senders))]

// 		msg := broker.Message{
// 			Topic:     topic,
// 			Sender:    sender,
// 			Payload:   utils.GeneratePayload(topic, sender),
// 			Timestamp: time.Now(),
// 		}

// 		// Measure publish latency
// 		start := time.Now()
// 		b.Publish(topic, msg)
// 		latency := time.Since(start).Microseconds()
// 		// Record metrics
// 		m.RecordMessage(topic, latency, msg)

// 		time.Sleep(time.Millisecond * 500)
// 	}
// }
