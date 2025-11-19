package dto

type PublishRequest struct {
	Topic   string                 `json:"topic" binding:"required"`
	Sender  string                 `json:"sender"`
	Payload map[string]interface{} `json:"payload"`
}

type SubscribeRequest struct {
	Topic string `form:"topic" binding:"required"`
}

// MetricsResponse represents the metrics response structure
type MetricsResponse struct {
	TotalMessages     int64                  `json:"totalMessages"`
	ActiveSubscribers int32                  `json:"activeSubscribers"`
	AvgLatency        int64                  `json:"avgLatency"`
	MessageRate       float64                `json:"messageRate"`
	Uptime            float64                `json:"uptime"`
	TopicMetrics      map[string]interface{} `json:"topicMetrics"`
}
