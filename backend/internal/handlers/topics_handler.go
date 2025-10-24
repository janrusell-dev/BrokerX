package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/janrusell-dev/brokerx/internal/broker"
)

func GetTopicsHandler(b *broker.Broker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"topics": b.GetTopics(),
		})
	}
}

func GetTopicInfoHandler(b *broker.Broker) gin.HandlerFunc {
	return func(c *gin.Context) {
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
	}
}

func GetAllTopicsInfoHandler(b *broker.Broker) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"topics": b.GetAllTopicsInfo(),
		})
	}
}
