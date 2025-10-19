package services

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/janrusell-dev/brokerx/internal/broker"
	"github.com/janrusell-dev/brokerx/internal/utils"
)

func StartSimulator(b *broker.Broker, m *MetricsService) {
	topics := []string{"alerts", "orders", "logs"}
	senders := []string{"service-A", "service-B", "service-C"}

	// Seed random numver generator
	utils.LogInfo("Message simulator started - publishing every 500ms")

	for {
		/// Pick random topic and sender
		topic := topics[rand.Intn(len(topics))]
		sender := senders[rand.Intn(len(senders))]

		msg := broker.Message{
			Topic:     topic,
			Sender:    sender,
			Payload:   generatePayload(topic, sender),
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

func generatePayload(topic, sender string) map[string]interface{} {
	switch topic {
	case "orders":
		return map[string]interface{}{
			"orderId":    fmt.Sprintf("ORD-%d", rand.Intn(100000)),
			"customerId": fmt.Sprintf("CUST-%d", rand.Intn(1000)),
			"amount":     rand.Float64() * 1000,
			"status":     []string{"pending", "processing", "completed", "cancelled"}[rand.Intn(4)],
			"items":      rand.Intn(10) + 1,
			"sender":     sender,
		}
	case "alerts":
		return map[string]interface{}{
			"alertId":  fmt.Sprintf("ALERT-%d", rand.Intn(1000)),
			"severity": []string{"low", "medium", "high", "critical"}[rand.Intn(4)],
			"message":  fmt.Sprintf("System alert from %s", sender),
			"source":   sender,
			"cpu":      rand.Intn(100),
			"memory":   rand.Intn(100),
		}
	case "logs":
		return map[string]interface{}{
			"logId":   fmt.Sprintf("LOG-%d", rand.Intn(100000)),
			"level":   []string{"debug", "info", "warn", "error"}[rand.Intn(4)],
			"message": fmt.Sprintf("Random log entry from %s", sender),
			"service": sender,
			"code":    rand.Intn(600),
		}
	default:
		return map[string]interface{}{
			"message": fmt.Sprintf("Random payload from %s", sender),
			"sender":  sender,
			"random":  rand.Intn(1000),
		}
	}

}
