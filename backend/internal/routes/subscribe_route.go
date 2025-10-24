package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/janrusell-dev/brokerx/internal/broker"
	"github.com/janrusell-dev/brokerx/internal/handlers"
	"github.com/janrusell-dev/brokerx/internal/services"
)

func RegisterSubscribeRoutes(r *gin.Engine, b *broker.Broker, m *services.MetricsService) {
	r.GET("/subscribe", handlers.SubscribeHandler(b, m))
}
