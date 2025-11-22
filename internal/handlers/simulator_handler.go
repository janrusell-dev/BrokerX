package handlers

import (
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/brokerx/internal/broker"
	"github.com/brokerx/internal/services"
	"github.com/brokerx/internal/utils"
	"github.com/gin-gonic/gin"
)

var (
	simulatorRunning bool
	simulatorDone    chan struct{}
	simulatorMu      sync.RWMutex
)

func StartSimulatorHandler(b *broker.Broker, m *services.MetricsService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		simulatorMu.Lock()
		if simulatorRunning {
			simulatorMu.Unlock()
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Simulator already running"})
			return
		}

		simulatorDone = make(chan struct{})
		simulatorRunning = true
		simulatorMu.Unlock()

		go func() {
			topics := []string{"orders", "alerts", "logs"}
			senders := []string{"service-A", "service-B", "service-C"}

			for {
				select {
				case <-simulatorDone:
					simulatorMu.Lock()
					simulatorRunning = false
					simulatorMu.Unlock()
					return
				default:
					for _, topic := range topics {
						sender := senders[rand.Intn(len(senders))]
						msg := broker.Message{
							Topic:     topic,
							Sender:    sender,
							Payload:   utils.GeneratePayload(topic, sender),
							Timestamp: time.Now(),
						}
						b.Publish(topic, msg)
						start := time.Now()
						latency := time.Since(start).Microseconds()
						m.RecordMessage(topic, latency, msg)
					}
					time.Sleep(500 * time.Millisecond)
				}
			}
		}()

		ctx.JSON(http.StatusOK, gin.H{"status": "Simulator started"})
	}
}

func StopSimulatorHandler(b *broker.Broker, m *services.MetricsService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		simulatorMu.Lock()
		defer simulatorMu.Unlock()

		if !simulatorRunning {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Simulator is not running"})
			return
		}

		close(simulatorDone)
		simulatorRunning = false

		ctx.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Simulator stopped successfully",
		})
	}
}

func SimulatorStatusHandler(b *broker.Broker, m *services.MetricsService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		simulatorMu.RLock()
		defer simulatorMu.RUnlock()
		ctx.JSON(http.StatusOK, gin.H{
			"running": simulatorRunning,
		})
	}
}
