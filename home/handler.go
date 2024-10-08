package home

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", HomeHandler)
}

func HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "home/home.jet", map[string]interface{}{
		"title": "Startseite",
	})
}
