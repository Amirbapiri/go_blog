package main

import (
	"fmt"
	"go_blog/internal/config"
	"go_blog/internal/handlers"
	"go_blog/internal/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

var app config.AppConfig

func main() {
	//We must create template cache only once.
	//We initialize template cache field that already created in configuration
	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can not create template cache", err)
	}
	app.TemplateCache = templateCache
	app.UseCache = false

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
