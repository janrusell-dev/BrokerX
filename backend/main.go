package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/janrusell-dev/brokerx/internal/broker"
	"github.com/janrusell-dev/brokerx/internal/handlers"
	"github.com/janrusell-dev/brokerx/internal/middleware"
	"github.com/janrusell-dev/brokerx/internal/services"
)

func main() {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	//initialize broker and metrics
	messageBroker := broker.NewBroker()
	metrics := services.NewMetricsService()

	handlers.RegisterPublishRoutes(r, messageBroker, metrics)
	handlers.RegisterSubscribeRoutes(r, messageBroker, metrics)
	handlers.RegisterMetricsRoutes(r, metrics)

	go services.StartSimulator(messageBroker, metrics)

	log.Printf("🚀 BrokerX backend running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
