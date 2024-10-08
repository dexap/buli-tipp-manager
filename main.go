package main

import (
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/username/buli-tipp-manager/auth"
	"github.com/username/buli-tipp-manager/home"
	// Importiere weitere Domänenpakete wie matches, odds
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Static files
	e.Static("/static", "static")

	// Jet Template Engine
	views := jet.NewSet(
		jet.NewOSFileSystemLoader("./"),
		jet.InDevelopmentMode(), // Entferne dies in der Produktion
	)

	// Renderer
	e.Renderer = &JetRenderer{
		views: views,
	}

	// Routes
	home.RegisterRoutes(e)
	auth.RegisterRoutes(e)
	// Registriere weitere Routen

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

type JetRenderer struct {
	views *jet.Set
}

func (r *JetRenderer) Render(w http.ResponseWriter, name string, data interface{}, c echo.Context) error {
	tmpl, err := r.views.GetTemplate(name)
	if err != nil {
		return err
	}

	var vars jet.VarMap
	if data != nil {
		vars = make(jet.VarMap)
		vars.Set("data", data)
	} else {
		vars = make(jet.VarMap)
	}

	return tmpl.Execute(w, vars, nil)
}
