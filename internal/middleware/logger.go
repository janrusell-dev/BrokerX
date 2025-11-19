package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery

		ctx.Next()

		duration := time.Since(start)
		statusCode := ctx.Writer.Status()

		// Color output based on status code
		statusColor := getStatusColor(statusCode)
		methodColor := getMethodColor(ctx.Request.Method)

		if query != "" {
			path = path + "?" + query
		}

		log.Printf("%s[%s]%s %s%-7s%s %s | %d | %s",
			methodColor, ctx.Request.Method, "\033[0m",
			statusColor, "", "\033[0m",
			path,
			statusCode,
			duration,
		)
	}
}

func getStatusColor(code int) string {
	switch {
	case code >= 200 && code < 300:
		return "\033[32m" // Green
	case code >= 300 && code < 400:
		return "\033[36m" // Cyan
	case code >= 400 && code < 500:
		return "\033[33m" // Yellow
	default:
		return "\033[31m" // Red
	}
}

func getMethodColor(method string) string {
	switch method {
	case "GET":
		return "\033[34m" // Blue
	case "POST":
		return "\033[36m" // Cyan
	case "PUT":
		return "\033[33m" // Yellow
	case "DELETE":
		return "\033[31m" // Red
	default:
		return "\033[37m" // White
	}
}
