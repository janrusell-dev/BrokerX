package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/janrusell-dev/brokerx/internal/services"
)

type MetricsHandler struct {
	Service *services.MetricsService
}

func NewMetricsHandler(s *services.MetricsService) *MetricsHandler {
	return &MetricsHandler{Service: s}
}

func (h *MetricsHandler) GetMetricsHandler(c *gin.Context) {
	metrics := h.Service.GetMetrics()
	uptime := time.Since(metrics.LastReset).Seconds()
	rate := h.Service.MessageRate()
	c.JSON(http.StatusOK, gin.H{
		"totalMessages":     metrics.TotalMessages,
		"avgLatency":        metrics.AvgLatency,
		"topicMetrics":      metrics.MessagePerTopic,
		"activeSubscribers": metrics.ActiveSubscribers,
		"latencyHistory":    metrics.LatencyHistory,
		"messageRate":       rate,
		"uptime":            uptime,
		"lastReset":         metrics.LastReset,
	})
}

// POST /metrics/reset
func (h *MetricsHandler) ResetMetricsHandler(c *gin.Context) {
	h.Service.ResetMetrics()
	c.JSON(http.StatusOK, gin.H{"status": "metrics reset"})
}
