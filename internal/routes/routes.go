package routes

import (
	"github.com/brokerx/internal/broker"
	"github.com/brokerx/internal/middleware"
	"github.com/brokerx/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupRouter(b *broker.Broker, m *services.MetricsService) *gin.Engine {
	r := gin.New()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	r.Use(
		middleware.Recovery(),
		middleware.RequestLogger(),
		middleware.CORSMiddleware(),
	)

	// Register all route groups
	RegisterMetricsRoutes(r, m)
	RegisterPublishRoutes(r, b, m)
	RegisterSubscribeRoutes(r, b, m)
	RegisterTopicRoutes(r, b)
	RegisterSimulator(r, b, m)

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "BrokerX",
		})
	})

	return r

}
