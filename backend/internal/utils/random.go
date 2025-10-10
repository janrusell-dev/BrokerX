package utils

import "math/rand"

func RandomTopic() string {
	topics := []string{"orders", "alerts", "payments", "logs"}
	return topics[rand.Intn(len(topics))]
}
