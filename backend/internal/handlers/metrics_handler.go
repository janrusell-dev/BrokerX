package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/janrusell-dev/brokerx/internal/services"
)

func RegisterMetricsRoutes(r *gin.Engine, m *services.MetricsService) {
	r.GET("/metrics", func(c *gin.Context) {
		c.JSON(http.StatusOK, m.GetMetrics())
	})
}
