package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/janrusell-dev/brokerx/internal/handlers"
	"github.com/janrusell-dev/brokerx/internal/services"
)

func RegisterMetricsRoutes(r *gin.Engine, s *services.MetricsService) {
	h := handlers.NewMetricsHandler(s)

	metricsGroup := r.Group("/metrics")
	metricsGroup.GET("", h.GetMetricsHandler)
	metricsGroup.POST("/reset", h.ResetMetricsHandler)
}
