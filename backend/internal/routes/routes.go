package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/janrusell-dev/brokerx/internal/broker"
	"github.com/janrusell-dev/brokerx/internal/middleware"
	"github.com/janrusell-dev/brokerx/internal/services"
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

	// Health check endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "BrokerX",
		})
	})

	return r

}
