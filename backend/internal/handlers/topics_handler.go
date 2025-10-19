package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/janrusell-dev/brokerx/internal/broker"
)

func RegisterTopicRoutes(r *gin.Engine, b *broker.Broker) {
	r.GET("/topics", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"topics": b.GetTopics(),
		})
	})

	r.GET("/topics/:topic", func(c *gin.Context) {
		topic := c.Param("topic")
		info := b.GetTopicInfo(topic)

		if !info["exists"].(bool) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Topic not found",
				"topic": topic,
			})
			return
		}

		c.JSON(http.StatusOK, info)
	})

	r.GET("/topics/info/all", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"topics": b.GetAllTopicsInfo(),
		})
	})
}
