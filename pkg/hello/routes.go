package hello

import (
	"github.com/labstack/echo"
)

// RegisterRoutes takes in an Echo router and registers routes onto it.
func RegisterRoutes(e *echo.Echo) {
	h := newHandler()

	e.GET("/hello", h.health)
}
