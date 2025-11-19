package routes

import (
	"github.com/brokerx/internal/handlers"
	"github.com/brokerx/internal/services"
	"github.com/gin-gonic/gin"
)

func RegisterMetricsRoutes(r *gin.Engine, s *services.MetricsService) {
	h := handlers.NewMetricsHandler(s)

	metricsGroup := r.Group("/metrics")
	metricsGroup.GET("", h.GetMetricsHandler)
	metricsGroup.POST("/reset", h.ResetMetricsHandler)
}
