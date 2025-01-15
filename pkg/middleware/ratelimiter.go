package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type RateLimiter struct {
	redisClient *redis.Client
	limit       int
	interval    time.Duration
}

func NewRateLimiter(redisClient *redis.Client, limit int, interval int) *RateLimiter {
	return &RateLimiter{
		redisClient: redisClient,
		limit:       limit,
		interval:    time.Duration(interval) * time.Second,
	}
}

func (rl *RateLimiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := "rate_limit:" + ip

		ctx := context.Background()
		count, err := rl.redisClient.Incr(ctx, key).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Rate limit error"})
			c.Abort()
			return
		}

		if count == 1 {
			rl.redisClient.Expire(ctx, key, rl.interval)
		}

		if count > int64(rl.limit) {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
