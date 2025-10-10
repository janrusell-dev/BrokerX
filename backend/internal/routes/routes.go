package routes

// import (
// 	"github.com/gin-gonic/gin"
// 	"github.com/janrusell-dev/brokerx/internal/middleware"
// )

// func initializeRoutes() {
// 	brokerRoutes := gin.Default()

// 	//middleware
// 	brokerRoutes.Use(middleware.CORSMiddleware())

// 	// Routes
// 	brokerRoutes.Use(middleware.Logger())
// 	brokerRoutes.POST("/publish", handlers.PublishHandler)
// 	brokerRoutes.Use(middleware.Logger())
// 	brokerRoutes.GET("/subscribe/:topic", handlers.SubscribeHandler)
// 	brokerRoutes.Use(middleware.Logger())
// 	brokerRoutes.GET("/stats", handlers.StatsHandler)

// 	// Start server
// 	brokerRoutes.Use(middleware.Logger())
// 	brokerRoutes.Run(":8080")
// }
