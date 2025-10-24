package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/janrusell-dev/brokerx/internal/broker"
	"github.com/janrusell-dev/brokerx/internal/handlers"
	"github.com/janrusell-dev/brokerx/internal/services"
)

func RegisterPublishRoutes(r *gin.Engine, b *broker.Broker, m *services.MetricsService) {
	r.POST("/publish", handlers.PublishMessageHandler(b, m))
}
