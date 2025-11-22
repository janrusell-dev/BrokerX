package routes

import (
	"github.com/brokerx/internal/broker"
	"github.com/brokerx/internal/handlers"
	"github.com/brokerx/internal/services"
	"github.com/gin-gonic/gin"
)

func RegisterSimulator(r *gin.Engine, b *broker.Broker, m *services.MetricsService) {
	s := r.Group("/simulator")
	s.POST("/start", handlers.StartSimulatorHandler(b, m))
	s.POST("/stop", handlers.StopSimulatorHandler(b, m))
	s.GET("/status", handlers.SimulatorStatusHandler(b, m))
}
