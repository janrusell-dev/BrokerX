package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/janrusell-dev/brokerx/internal/broker"
	"github.com/janrusell-dev/brokerx/internal/services"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func RegisterSubscribeRoutes(r *gin.Engine, b *broker.Broker, m *services.MetricsService) {
	r.GET("/subscribe", func(ctx *gin.Context) {
		topic := ctx.Query("topic")
		if topic == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing topic"})
			return
		}
		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			return
		}
		defer conn.Close()

		ch := b.Subscribe(topic)
		m.IncrementSubscribers()

		for msg := range ch {
			if err := conn.WriteJSON(msg); err != nil {
				break
			}
		}

		m.DecrementSubscribers()
	})
}
