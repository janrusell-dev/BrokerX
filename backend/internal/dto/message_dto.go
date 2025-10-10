package dto

type PublishRequest struct {
	Topic   string `json:"topic" binding:"required"`
	Sender  string `json:"sender" binding:"required"`
	Payload string `json:"payload" binding:"required"`
}
