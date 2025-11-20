package routes

import (
	"github.com/brokerx/internal/broker"
	"github.com/brokerx/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterTopicRoutes(r *gin.Engine, b *broker.Broker) {
	topicsGroup := r.Group("/topics")
	topicsGroup.GET("/topics", handlers.GetTopicsHandler(b))
	topicsGroup.GET("/topics/:topic", handlers.GetTopicInfoHandler(b))
	topicsGroup.GET("/topics/info/all", handlers.GetAllTopicsInfoHandler(b))
}
