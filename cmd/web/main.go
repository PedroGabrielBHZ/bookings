package main

import (
	"app/pkg/config"
	"app/pkg/handlers"
	"time"

	"fmt"
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

const port = ":8080"

var app config.AppConfig

var session *scs.SessionManager

func StartServer() {
	app.InProduction = false // Set to true in production
	initializeSession()

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	fmt.Println("Access the server at: http://localhost" + port)
	fmt.Println("Press Ctrl+C to stop the server")
	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err := srv.ListenAndServe()
	log.Fatal(err)
}

func initializeSession() {
	session = scs.New()
	session.Lifetime = 24 * time.Hour // 24 hours
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // Set to true in production
	app.Session = session
}

func main() {
	StartServer()
}
