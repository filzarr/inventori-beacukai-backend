package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func AccessLog(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()
	if err != nil {
		return err
	}
	log.Info().
		Str("method", c.Method()).
		Str("path", c.Path()).
		Any("query", c.Queries()).
		Str("ip", c.IP()).
		Str("user_agent", c.Get("User-Agent")).
		Dur("duration", time.Since(start)). // duration in ms
		Msg("access log")
	return nil
}
