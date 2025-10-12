package utils

import (
	"fmt"
	"log"
	"time"
)

const (
	ColorReset   = "\033[0m"
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorMagenta = "\033[35m"
	ColorCyan    = "\033[36m"
	ColorWhite   = "\033[37m"
)

// LogEvent logs a formatted event with timestamp (cyan)
func LogEvent(message string) {
	timestamp := time.Now().Format("15:04:05")
	log.Printf("%s[%s]%s %s", ColorCyan, timestamp, ColorReset, message)
}

// LogSuccess logs a success message in green with checkmark
func LogSuccess(message string) {
	timestamp := time.Now().Format("15:04:05")
	log.Printf("%s✓ [%s]%s %s", ColorGreen, timestamp, ColorReset, message)
}

// LogError logs an error message in red with X mark
func LogError(message string, err error) {
	timestamp := time.Now().Format("15:04:05")
	if err != nil {
		log.Printf("%s✗ [%s]%s %s: %v", ColorRed, timestamp, ColorReset, message, err)
	} else {
		log.Printf("%s✗ [%s]%s %s", ColorRed, timestamp, ColorReset, message)
	}
}

// LogWarning logs a warning message in yellow with warning symbol
func LogWarning(message string) {
	timestamp := time.Now().Format("15:04:05")
	log.Printf("%s⚠ [%s]%s %s", ColorYellow, timestamp, ColorReset, message)
}

// LogInfo logs an info message in blue with info symbol
func LogInfo(format string, args ...interface{}) {
	timestamp := time.Now().Format("15:04:05")
	message := fmt.Sprintf(format, args...)
	log.Printf("%sℹ [%s]%s %s", ColorBlue, timestamp, ColorReset, message)
}

// LogDebug logs a debug message in magenta
func LogDebug(format string, args ...interface{}) {
	timestamp := time.Now().Format("15:04:05")
	message := fmt.Sprintf(format, args...)
	log.Printf("%s🔍 [%s]%s %s", ColorMagenta, timestamp, ColorReset, message)
}

// LogRequest logs an HTTP request (used for custom logging)
func LogRequest(method, path string, statusCode int, duration time.Duration) {
	timestamp := time.Now().Format("15:04:05")
	statusColor := getStatusColor(statusCode)
	methodColor := getMethodColor(method)

	log.Printf("%s[%s]%s %s%-7s%s %s | %s%d%s | %v",
		ColorCyan, timestamp, ColorReset,
		methodColor, method, ColorReset,
		path,
		statusColor, statusCode, ColorReset,
		duration,
	)
}

// LogPublish logs a message publish event
func LogPublish(topic string, sender string, latency int64) {
	timestamp := time.Now().Format("15:04:05")
	log.Printf("%s📤 [%s]%s Published to %s%s%s from %s%s%s (latency: %dms)",
		ColorGreen, timestamp, ColorReset,
		ColorCyan, topic, ColorReset,
		ColorYellow, sender, ColorReset,
		latency,
	)
}

// LogSubscribe logs a subscription event
func LogSubscribe(topic string, subscriberCount int) {
	timestamp := time.Now().Format("15:04:05")
	log.Printf("%s📥 [%s]%s New subscriber to %s%s%s (total: %d)",
		ColorGreen, timestamp, ColorReset,
		ColorCyan, topic, ColorReset,
		subscriberCount,
	)
}

// LogUnsubscribe logs an unsubscription event
func LogUnsubscribe(topic string, subscriberCount int) {
	timestamp := time.Now().Format("15:04:05")
	log.Printf("%s📤 [%s]%s Subscriber left %s%s%s (remaining: %d)",
		ColorYellow, timestamp, ColorReset,
		ColorCyan, topic, ColorReset,
		subscriberCount,
	)
}

// getStatusColor returns color based on HTTP status code
func getStatusColor(code int) string {
	switch {
	case code >= 200 && code < 300:
		return ColorGreen
	case code >= 300 && code < 400:
		return ColorCyan
	case code >= 400 && code < 500:
		return ColorYellow
	default:
		return ColorRed
	}
}

// getMethodColor returns color based on HTTP method
func getMethodColor(method string) string {
	switch method {
	case "GET":
		return ColorBlue
	case "POST":
		return ColorCyan
	case "PUT":
		return ColorYellow
	case "DELETE":
		return ColorRed
	case "PATCH":
		return ColorMagenta
	default:
		return ColorWhite
	}
}
