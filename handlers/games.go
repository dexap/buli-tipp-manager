package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GamesHandler(c echo.Context) error {
	games := []struct {
		TeamA string
		TeamB string
		Odds  float64
	}{
		{"Team A", "Team B", 1.5},
		{"Team C", "Team D", 2.0},
	}

	data := map[string]interface{}{
		"Title":  "Spiele",
		"Games":  games,
		"IsHTMX": c.Request().Header.Get("HX-Request") == "true",
	}
	// Überprüfen, ob es eine htmx-Anfrage ist
	if data["IsHTMX"].(bool) {
		// htmx-Anfrage: nur den Content rendern
		return c.Render(http.StatusOK, "content", data)
	} else {
		// Normale Anfrage: gesamtes Template rendern
		return c.Render(http.StatusOK, "games", data)
	}
}
