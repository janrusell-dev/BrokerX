package routes

import (
	"github.com/brokerx/internal/broker"
	"github.com/brokerx/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterTopicRoutes(r *gin.Engine, b *broker.Broker) {
	r.GET("/topics", handlers.GetTopicsHandler(b))
	r.GET("/topics/:topic", handlers.GetTopicInfoHandler(b))
	r.GET("/topics/info/all", handlers.GetAllTopicsInfoHandler(b))
}
