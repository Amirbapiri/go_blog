package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"go_blog/internal/config"
	"go_blog/internal/handlers"
	"go_blog/internal/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig

var sessionManager *scs.SessionManager

func main() {
	//We must create template cache only once.
	//We initialize template cache field that already created in configuration
	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can not create template cache", err)
	}
	app.TemplateCache = templateCache
	app.UseCache = false

	//We do use this flag to control some configurations.
	//Surely, needs to be changed in production mode.
	app.InProduction = false

	//Setting up session manager
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = app.InProduction

	app.Session = sessionManager

	repository := handlers.CreateNewRepository(&app)
	handlers.SetRepository(repository)

	render.SetConfigForTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = serve.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
