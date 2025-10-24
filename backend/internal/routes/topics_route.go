package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/janrusell-dev/brokerx/internal/broker"
	"github.com/janrusell-dev/brokerx/internal/handlers"
)

func RegisterTopicRoutes(r *gin.Engine, b *broker.Broker) {
	r.GET("/topics", handlers.GetTopicsHandler(b))
	r.GET("/topics/:topic", handlers.GetTopicInfoHandler(b))
	r.GET("/topics/info/all", handlers.GetAllTopicsInfoHandler(b))
}
