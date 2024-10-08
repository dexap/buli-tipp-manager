package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func IndexHandler(c echo.Context) error {
	data := map[string]interface{}{
		"Title": "Startseite",
	}

	// Prüfen, ob die Anfrage von htmx kommt
	if c.Request().Header.Get("HX-Request") == "true" {
		// htmx-Anfrage: nur den Inhalt zurückgeben
		return c.Render(http.StatusOK, "index_partial.html", data)
	} else {
		// Normale Anfrage: das gesamte Layout zurückgeben
		return c.Render(http.StatusOK, "index.html", data)
	}
}
