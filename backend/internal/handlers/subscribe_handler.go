package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/janrusell-dev/brokerx/internal/broker"
	"github.com/janrusell-dev/brokerx/internal/services"
	"github.com/janrusell-dev/brokerx/internal/utils"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func RegisterSubscribeRoutes(r *gin.Engine, b *broker.Broker, m *services.MetricsService) {
	r.GET("/subscribe", func(ctx *gin.Context) {
		topic := ctx.Query("topic")
		if topic == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing topic"})
			return
		}

		utils.LogInfo("WebSocket connection request for topic: %s from %s", topic, ctx.ClientIP())
		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			utils.LogError("WebSocket upgrade failed", err)
			return
		}
		defer conn.Close()

		ch := b.Subscribe(topic)
		defer b.Unsubscribe(topic, ch)
		m.IncrementSubscribers()
		defer m.DecrementSubscribers()

		utils.LogSuccess("WebSocket connected to topic: %s")

		// Send connection confirmation
		if err := conn.WriteJSON(gin.H{
			"type":    "connected",
			"topic":   topic,
			"message": "Successfully subscribed",
		}); err != nil {
			utils.LogError("Failed to send connection confirmation", err)
			return
		}

		// Listen for messages from the channel
		for msg := range ch {
			if err := conn.WriteJSON(msg); err != nil {
				utils.LogWarning("Failed to send message to subscriber, closing connection")
				break
			}
		}

		utils.LogInfo("WebSocket disconnected from topic: %s", topic)
	})
}
