package middleware

import (
	"log"

	"github.com/actanonvebra/honeyshop/internal/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RateLimiterMiddleWare(logService services.LogService) echo.MiddlewareFunc {
	return middleware.RateLimiterWithConfig(middleware.RateLimiterConfig{
		Store: middleware.NewRateLimiterMemoryStore(10),
		IdentifierExtractor: func(c echo.Context) (string, error) {
			return c.RealIP(), nil
		},
		ErrorHandler: func(c echo.Context, err error) error {
			ip := c.RealIP()
			log.Printf("Rate limit exceeded for IP: %s", ip)
			logService.LogAttack("Brute Force", "Too many login attempts", ip)
			return echo.ErrTooManyRequests
		},
	})
}
