package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/bciupka/go-room-booking/pkg/config"
	"github.com/bciupka/go-room-booking/pkg/handlers"
	"github.com/bciupka/go-room-booking/pkg/render"
)

const port = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	tc, err := render.CreateCacheForTemplates()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	fmt.Println("Server starting on port", port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
