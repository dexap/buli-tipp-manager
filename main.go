package main

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	"github.com/dexap/buli-tipp-manager/templates"
)

func main() {
	e := echo.New()

	// Statische Dateien bereitstellen
	e.Static("/static", "static")

	// Route für die Startseite
	e.GET("/", func(c echo.Context) error {
		// Render das Template
		return templ.Execute(c.Response().Writer, c.Request().Context(), templates.IndexPage())
	})

	// Server starten
	e.Logger.Fatal(e.Start(":8080"))
}
