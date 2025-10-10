package utils

import (
	"log"
	"time"
)

func LogEvent(event string) {
	log.Printf("[%s] %s\n", time.Now().Format("15:04:05"), event)
}
