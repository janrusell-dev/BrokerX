package main

import (
	"log"

	"github.com/brokerx/internal/broker"
	"github.com/brokerx/internal/routes"
	"github.com/brokerx/internal/services"
	"github.com/brokerx/internal/utils"
)

func main() {
	utils.LogInfo("Initializing BrokerX...")

	//Initialize core services
	messageBroker := broker.NewBroker()
	metrics := services.NewMetricsService(1000)

	utils.LogSuccess("Broker and metrics service initialized")

	// Setup router with all handlers
	r := routes.SetupRouter(messageBroker, metrics)

	// Start message simulator in background
	go services.StartSimulator(messageBroker, metrics)
	utils.LogInfo("Message simulator started")

	// Print startup banner
	printBanner()

	utils.LogSuccess("BrokerX backend running at http://localhost:8080")

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func printBanner() {
	banner := `
╔══════════════════════════════════════════════════════╗
║                                                      ║
║        🚀 BrokerX Message Broker v1.0                ║
║                                                      ║
║        A lightweight, real-time message broker       ║
║                                                      ║
╚══════════════════════════════════════════════════════╝

📡 Available Endpoints:
   GET  /health              - Health check
   POST /publish             - Publish message to topic
   GET  /subscribe?topic=x   - Subscribe to topic (WebSocket)
   GET  /metrics             - Get system metrics
   POST /metrics/reset       - Reset metrics
   GET  /metrics/latency     - Get latency history
   GET  /topics              - List all active topics
   GET  /topics/:topic       - Get topic information
   GET  /topics/info/all     - Get all topics info

`
	log.Print(banner)
}
