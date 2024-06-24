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

func RateLimiterFiber(max_requests int, duration time.Duration) fiber.Handler {
	// Create a rate limiter using the middleware
	return limiter.New(limiter.Config{
		Max:        max_requests,
		Expiration: duration,
		KeyGenerator: func(context *fiber.Ctx) string {
			// Use the request IP address as the key for rate limiting
			return context.IP()
		},
		LimitReached: func(context *fiber.Ctx) error {
			return controllers.ErrorHandlerFiber(context, http.StatusTooManyRequests, "Rate limit exceeded", "RateLimiterFiber | RateLimiter")
		},
	})
}

func RateLimiterGin(max_requests int, duration time.Duration) gin.HandlerFunc {
	// Create a rate limiter using token bucket algorithm
	limiter := ratelimit.NewBucketWithQuantum(duration, int64(max_requests), int64(max_requests))

	return func(context *gin.Context) {
		// Check if the request IP address exceeds the rate limit
		if limiter.TakeAvailable(1) < 1 {
			controllers.ErrorHandlerGin(context, http.StatusTooManyRequests, "Rate limit exceeded", "RateLimiterGin | RateLimiter")
			context.Abort()
			return
		}

		context.Next()
	}
}
