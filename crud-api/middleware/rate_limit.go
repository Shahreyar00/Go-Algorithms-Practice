package middleware

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	clients     = make(map[string][]time.Time) // Track all request times for each IP
	clientsLock = &sync.Mutex{}
	rateLimit   = 10
	timeWindow  = time.Minute // per minute
)

func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		now := time.Now()

		clientsLock.Lock()
		defer clientsLock.Unlock()

		// Clean up old requests outside the time window
		requests, exists := clients[clientIP]
		if exists {
			validRequests := make([]time.Time, 0, len(requests))
			for _, timestamp := range requests {
				if now.Sub(timestamp) <= timeWindow {
					validRequests = append(validRequests, timestamp)
				}
			}
			clients[clientIP] = validRequests
		} else {
			clients[clientIP] = []time.Time{}
		}

		// Check if the rate limit is exceeded
		if len(clients[clientIP]) >= rateLimit {
			log.Printf("Rate limit exceeded for IP: %s at %s", clientIP, now.Format(time.RFC3339))
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}

		// Register the request
		clients[clientIP] = append(clients[clientIP], now)
		c.Next()
	}
}
