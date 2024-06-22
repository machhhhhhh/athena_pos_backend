package middlewares

import (
	"athena-pos-backend/controllers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/juju/ratelimit"
)

func RateLimiterFiber(maxRequests int, duration time.Duration) fiber.Handler {
	// Create a rate limiter using the middleware
	return limiter.New(limiter.Config{
		Max:        maxRequests,
		Expiration: duration,
		KeyGenerator: func(c *fiber.Ctx) string {
			// Use the request IP address as the key for rate limiting
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error":         "Rate limit exceeded",
				"error_section": "RateLimiter",
			})
		},
	})
}

func RateLimiterGin(max_requests int, duration time.Duration) gin.HandlerFunc {
	// Create a rate limiter using token bucket algorithm
	limiter := ratelimit.NewBucketWithQuantum(duration, int64(max_requests), int64(max_requests))

	return func(c *gin.Context) {
		// Check if the request IP address exceeds the rate limit
		if limiter.TakeAvailable(1) < 1 {
			controllers.ErrorHandlerGin(c, http.StatusTooManyRequests, "Rate limit exceeded", "RateLimiter")
			c.Abort()
			return
		}

		c.Next()
	}
}
