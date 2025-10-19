package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/janrusell-dev/brokerx/internal/broker"
	"github.com/janrusell-dev/brokerx/internal/dto"
	"github.com/janrusell-dev/brokerx/internal/services"
	"github.com/janrusell-dev/brokerx/internal/utils"
)

func RegisterPublishRoutes(r *gin.Engine, b *broker.Broker, m *services.MetricsService) {
	r.POST("/publish", func(ctx *gin.Context) {
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

		m.RecordMessage(req.Topic, latency)
		utils.LogEvent("Published message to topic: " + req.Topic)

		ctx.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Message published",
			"topic":   req.Topic,
			"latency": latency,
		})

	})
}
