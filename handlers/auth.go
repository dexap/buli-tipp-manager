package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var users = map[string]string{}

// AuthHandler rendert die Login-/Signup-Seite und verarbeitet Formulareingaben
func AuthHandler(c echo.Context) error {
	if c.Request().Method == http.MethodGet {
		// GET-Anfrage: Zeige das Formular an
		return c.Render(http.StatusOK, "auth.html", nil)
	} else if c.Request().Method == http.MethodPost {
		action := c.FormValue("action")
		email := c.FormValue("email")
		password := c.FormValue("password")

		if action == "continue" {
			// Überprüfe, ob der Benutzer existiert
			if _, exists := users[email]; exists {
				// Benutzer existiert, zeige Login-Button
				return c.HTML(http.StatusOK,
					` <p>Anmeldung fehlgeschlagen. Bitte versuchen Sie es erneut oder registrieren Sie sich.</p>
					<div class="space-x-2">
						<button class="btn btn-primary"
										hx-post="/auth"
										hx-vals='{"action":"login","email":"`+email+`","password":"`+password+`"}'>Erneut anmelden</button>
						<button class="btn btn-secondary"
										hx-post="/auth"
										hx-vals='{"action":"signup","email":"`+email+`","password":"`+password+`"}'>Registrieren</button>
					</div>
					`)
			} else {
				// Benutzer existiert nicht, biete Registrierung an
				return c.HTML(http.StatusOK,
					`	<p>Bitte lösen Sie das Captcha, um sich zu registrieren.</p>
					<form hx-post="/auth" hx-target="#auth-area" hx-swap="innerHTML">
						<input type="hidden" name="action" value="signup">
						<input type="hidden" name="email" value="`+email+`">
						<input type="hidden" name="password" value="`+password+`">
						<!-- Hier würde das Captcha eingebunden -->
						<p>Was ist 2 + 2?</p>
						<input type="text" name="captcha" placeholder="Antwort eingeben" required>
						<button type="submit">Registrieren</button>
					</form>`)
			}
		} else if action == "login" {
			// Login-Logik
			if storedPassword, exists := users[email]; exists && storedPassword == password {
				return c.String(http.StatusOK, "Anmeldung erfolgreich für "+email)
			}
			return c.String(http.StatusOK, "Anmeldedaten falsch. Bitte erneut versuchen.")
		} else if action == "signup" {
			// Signup-Logik mit Captcha
			captcha := c.FormValue("captcha")
			if !verifyCaptcha(captcha) {
				return c.String(http.StatusOK, "Captcha ungültig.")
			}
			users[email] = password
			return c.String(http.StatusOK, "Registrierung erfolgreich für "+email)
		}
	}
	return c.NoContent(http.StatusBadRequest)
}

// Funktion zur Captcha-Überprüfung (Dummy-Funktion)
func verifyCaptcha(captcha string) bool {
	// Hier die tatsächliche Captcha-Überprüfung einfügen
	return captcha == "1234" // Platzhalter für Demo-Zwecke
}
