package tasks

import (
	"net/http"

	"github.com/apex/log"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var ip_limiter = make(map[string]*rate.Limiter)

func Task1() {
	router := gin.Default()

	router.POST("/api/post", func(c *gin.Context) {
		ip := c.ClientIP()
		log.Debug(ip)
		limiter := getLimiter(ip)

		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate Limit exceeded!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "POST reqest successful", "limiter": limiter, "ip": ip})
	})

	if err := router.Run(":8080"); err != nil {
		log.WithError(err).Error("Failed to listen")
	}
	log.Info("Server is listening on port: 8080")
}

func getLimiter(ip string) *rate.Limiter {
	limiter, exists := ip_limiter[ip]
	if !exists {
		limiter = rate.NewLimiter(1, 1)
		ip_limiter[ip] = limiter
	}

	return limiter
}
