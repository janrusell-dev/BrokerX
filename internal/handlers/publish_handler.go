package handlers

import (
	"net/http"
	"time"

	"github.com/brokerx/internal/broker"
	"github.com/brokerx/internal/dto"
	"github.com/brokerx/internal/services"
	"github.com/brokerx/internal/utils"
	"github.com/gin-gonic/gin"
)

func PublishMessageHandler(b *broker.Broker, m *services.MetricsService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.PublishRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		msg := broker.Message{
			Topic:     req.Topic,
			Sender:    req.Sender,
			Payload:   req.Payload,
			Timestamp: time.Now(),
		}

		start := time.Now()
		b.Publish(req.Topic, msg)
		latency := time.Since(start).Microseconds()

		m.RecordMessage(req.Topic, latency, msg)
		utils.LogEvent("Published message to topic: " + req.Topic)

		ctx.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Message published",
			"topic":   req.Topic,
			"latency": latency,
		})

	}
}
