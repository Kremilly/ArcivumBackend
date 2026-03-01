package middleware

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
)

func LoginLimiter() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        5,
		Expiration: 1 * time.Minute,

		KeyGenerator: func(c fiber.Ctx) string {
			return c.IP()
		},

		LimitReached: func(c fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"success": false,
				"error":   "Too many login attempts. Please wait 1 minute.",
			})
		},

		Next: func(c fiber.Ctx) bool {
			return c.IP() == "127.0.0.1" || c.IP() == "::1"
		},
	})
}
