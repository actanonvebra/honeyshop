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
			ip := c.RealIP()
			log.Printf("Ratge ")
			log.Printf("Rate limit exceeded for IP: %s", ip)
			err := logService.LogAttack("Brute Force", "Too many login attempts", ip)
			if err != nil {
				log.Printf("LogAttack error: %v", err)
			} else {
				log.Println("LogAttack success")
			}
			return ip, nil
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return echo.ErrTooManyRequests
		},
	})
}
