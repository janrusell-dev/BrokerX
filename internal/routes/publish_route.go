package routes

import (
	"github.com/brokerx/internal/broker"
	"github.com/brokerx/internal/handlers"
	"github.com/brokerx/internal/services"
	"github.com/gin-gonic/gin"
)

func RegisterPublishRoutes(r *gin.Engine, b *broker.Broker, m *services.MetricsService) {
	r.POST("/publish", handlers.PublishMessageHandler(b, m))
}
