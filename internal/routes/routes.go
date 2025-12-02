package routes

import (
	"log"

	"github.com/brokerx/internal/broker"
	"github.com/brokerx/internal/middleware"
	"github.com/brokerx/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupRouter(b *broker.Broker, m *services.MetricsService) *gin.Engine {
	r := gin.New()
	if err := r.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Printf("warning: failed to set trusted proxies %v", err)
	}

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
