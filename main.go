package main

import (
	"html/template"
	"io"

	"github.com/dexap/buli-tipp-manager/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	tmpl := template.New("templates")
	// Template-Renderer einrichten
	tmpl, err := tmpl.ParseFiles(
		"templates/base.html")

	if err != nil {
		e.Logger.Fatal(err)
	}

	_, err = tmpl.ParseGlob("templates/*.html")
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Renderer = &TemplateRenderer{
		templates: tmpl,
	}

	// Route für die Startseite
	e.GET("/", handlers.IndexHandler)
	e.GET("/games", handlers.GamesHandler)
	e.GET("/auth", handlers.AuthHandler)
	e.POST("/auth", handlers.AuthHandler)

	// Statische Dateien bereitstellen
	e.Static("/static", "static")

	// Server starten
	e.Logger.Fatal(e.Start(":8080"))
}

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Überprüfen, ob es eine htmx-Anfrage ist
	if d, ok := data.(map[string]interface{}); ok {
		if isHTMX, ok := d["IsHTMX"].(bool); ok && isHTMX && name != "content" {
			// htmx-Anfrage: nur den Content rendern
			return t.templates.ExecuteTemplate(w, "content", data)
		}
	}
	return t.templates.ExecuteTemplate(w, name, data)
}
