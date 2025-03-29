package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf creates CSRF protecion for all POSTs to a server
func NoSurf(next http.Handler) http.Handler {
	csrfToken := nosurf.New(next)
	csrfToken.SetBaseCookie(http.Cookie{
		Path:     "/",
		Secure:   app.InProduction,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfToken
}

// SessionLoad loads a session every time the page is being reloaded
func SessionLoad(next http.Handler) http.Handler {
	return app.Session.LoadAndSave(next)
}
