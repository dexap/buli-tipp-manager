package auth

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/auth/login", LoginGetHandler)
	e.POST("/auth/login", LoginPostHandler)
}

func LoginGetHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "auth/login.jet", map[string]interface{}{
		"title": "Login",
	})
}

func LoginPostHandler(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	// Hier würdest du die Benutzerauthentifizierung implementieren
	// Für dieses Beispiel leiten wir einfach zur Startseite weiter
	fmt.Println(email, HashPassword(password))

	return c.Redirect(http.StatusSeeOther, "/")
}

func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(fmt.Errorf("error hashing password: %v", err))
	}
	return string(hash)
}
