package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/janrusell-dev/brokerx/internal/broker"
	"github.com/janrusell-dev/brokerx/internal/handlers"
	"github.com/janrusell-dev/brokerx/internal/middleware"
	"github.com/janrusell-dev/brokerx/internal/services"
)

func SetupRouter(b *broker.Broker, m *services.MetricsService) *gin.Engine {
	r := gin.New()

	r.Use(
		middleware.Recovery(),
		middleware.RequestLogger(),
		middleware.CORSMiddleware(),
	)

	// Register all route groups
	handlers.RegisterMetricsRoutes(r, m)
	handlers.RegisterPublishRoutes(r, b, m)
	handlers.RegisterSubscribeRoutes(r, b, m)
	handlers.RegisterTopicRoutes(r, b)

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "BrokerX",
		})
	})

	return r

}
