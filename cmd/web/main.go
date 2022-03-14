package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/pedrogabrielbhz/bookings/pkg/config"
	"github.com/pedrogabrielbhz/bookings/pkg/handlers"
	"github.com/pedrogabrielbhz/bookings/pkg/render"

	"github.com/alexedwards/scs/v2"
)

var app config.AppConfig

var session *scs.SessionManager

const portNumber = ":8080"

// main is the main function
func main() {
	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf(
		"Starting application on port http://localhost%s... \n",
		portNumber,
	)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
