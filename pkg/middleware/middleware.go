package middleware

import (
	"clean-arch/config"
	"github.com/labstack/echo"
)

type GoMiddleware struct {
	config *config.Config
}

// CORS will handle the CORS middleware
func (m *GoMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(c)
	}
}

// InitMiddleware initialize the middleware
func InitMiddleware(config *config.Config) *GoMiddleware {
	return &GoMiddleware{
		config: config,
	}
}

func (m *GoMiddleware) Auth(authorization bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	}
}
