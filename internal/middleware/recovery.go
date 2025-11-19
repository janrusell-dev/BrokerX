package middleware

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("\033[31m⚠️  PANIC RECOVERED\033[0m")
				log.Printf("Error: %v", err)
				log.Printf("Stack trace:\n%s", debug.Stack())

				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error":  "Internal server error",
					"detail": fmt.Sprintf("%v", err),
				})
			}
		}()
		ctx.Next()
	}
}
